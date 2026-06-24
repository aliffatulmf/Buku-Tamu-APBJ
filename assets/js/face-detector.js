import { FaceDetector, FilesetResolver } from "https://cdn.jsdelivr.net/npm/@mediapipe/tasks-vision@0.10.18/vision_bundle.mjs";

let faceDetector = null;
let isDetecting = false;
let lastVideoTime = -1;
let overlayCanvas = null;
let overlayCtx = null;
let detectionResults = [];
let initPromise = null;

async function initFaceDetector() {
    if (initPromise) return initPromise;

    initPromise = (async () => {
        try {
            console.log("Initializing face detector...");

            const vision = await FilesetResolver.forVisionTasks("https://cdn.jsdelivr.net/npm/@mediapipe/tasks-vision@0.10.18/wasm");
            console.log("FilesetResolver loaded");

            faceDetector = await FaceDetector.createFromOptions(vision, {
                baseOptions: {
                    modelAssetPath:
                        "https://storage.googleapis.com/mediapipe-models/face_detector/blaze_face_short_range/float16/latest/blaze_face_short_range.tflite",
                    delegate: "CPU",
                },
                runningMode: "VIDEO",
                minDetectionConfidence: 0.5,
                minSuppressionThreshold: 0.3,
            });

            console.log("Face detector created successfully");
            return true;
        } catch (error) {
            console.error("Failed to init face detector:", error);
            faceDetector = null;
            initPromise = null;
            return false;
        }
    })();

    return initPromise;
}

function setupOverlayCanvas(videoElement) {
    console.log("Setting up overlay canvas...");

    const wrapper = videoElement.closest(".webcam-preview-wrapper");
    if (!wrapper) {
        console.error("Could not find webcam-preview-wrapper");
        return null;
    }

    const existing = wrapper.querySelector(".face-overlay");
    if (existing) existing.remove();

    overlayCanvas = document.createElement("canvas");
    overlayCanvas.className = "face-overlay";
    overlayCanvas.style.cssText = `
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        pointer-events: none;
        border-radius: 12px;
        z-index: 10;
    `;

    wrapper.appendChild(overlayCanvas);
    overlayCtx = overlayCanvas.getContext("2d");

    const updateSize = () => {
        if (videoElement.videoWidth > 0) {
            overlayCanvas.width = videoElement.videoWidth;
            overlayCanvas.height = videoElement.videoHeight;
            console.log(
                "Canvas size set:",
                videoElement.videoWidth,
                "x",
                videoElement.videoHeight,
            );
        } else {
            requestAnimationFrame(updateSize);
        }
    };
    updateSize();

    return overlayCanvas;
}

function startDetectionLoop(videoElement) {
    if (!faceDetector) {
        console.error("Face detector not initialized");
        return;
    }

    if (isDetecting) {
        console.log("Detection already running");
        return;
    }

    console.log("Starting detection loop...");
    isDetecting = true;
    lastVideoTime = -1;

    function detect() {
        if (
            !isDetecting ||
            !videoElement ||
            videoElement.paused ||
            videoElement.ended
        ) {
            return;
        }

        if (videoElement.readyState < 2) {
            requestAnimationFrame(detect);
            return;
        }

        if (videoElement.currentTime !== lastVideoTime) {
            lastVideoTime = videoElement.currentTime;

            try {
                const result = faceDetector.detectForVideo(
                    videoElement,
                    performance.now(),
                );
                detectionResults = result.detections || [];

                if (detectionResults.length > 0) {
                    console.log("Faces detected:", detectionResults.length);
                }

                drawBoundingBoxes();
            } catch (e) {
                console.error("Detection error:", e);
            }
        }

        requestAnimationFrame(detect);
    }

    detect();
}

function stopDetectionLoop() {
    console.log("Stopping detection loop...");
    isDetecting = false;
    detectionResults = [];

    if (overlayCtx && overlayCanvas) {
        overlayCtx.clearRect(0, 0, overlayCanvas.width, overlayCanvas.height);
    }
}

function drawBoundingBoxes() {
    if (!overlayCtx || !overlayCanvas) return;

    const video = document.getElementById("camera");
    if (!video) return;

    if (video.videoWidth > 0 && overlayCanvas.width !== video.videoWidth) {
        overlayCanvas.width = video.videoWidth;
        overlayCanvas.height = video.videoHeight;
    }

    overlayCtx.clearRect(0, 0, overlayCanvas.width, overlayCanvas.height);

    detectionResults.forEach((detection) => {
        const bbox = detection.boundingBox;

        const flippedX = overlayCanvas.width - bbox.originX - bbox.width;

        overlayCtx.strokeStyle = "#4361ee";
        overlayCtx.lineWidth = 3;
        overlayCtx.strokeRect(flippedX, bbox.originY, bbox.width, bbox.height);
    });
}

function getFaceCount() {
    return detectionResults.length;
}

function validateFaceCount() {
    const count = getFaceCount();

    if (count === 0) {
        return { valid: false, message: "Tidak ada wajah terdeteksi" };
    }

    if (count > 2) {
        return { valid: false, message: "Terlalu banyak wajah, maksimal 2 orang" };
    }

    return { valid: true, message: `${count} wajah terdeteksi` };
}

window.FaceDetection = {
    init: initFaceDetector,
    setupOverlay: setupOverlayCanvas,
    startLoop: startDetectionLoop,
    stopLoop: stopDetectionLoop,
    validate: validateFaceCount,
    getCount: getFaceCount,
};

console.log("Face detection module loaded");
