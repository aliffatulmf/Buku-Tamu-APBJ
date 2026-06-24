// Registrasi Common JavaScript
// Shared functions for Pemda and Penyedia registration pages

// ============================================
// Face Detection State
// ============================================
var faceDetectorReady = false;

// ============================================
// Webcam Functions
// ============================================
var webcamInstance = null;

function initWebcam() {
    const camera = document.getElementById("camera")
    const canvas = document.getElementById("result")
    webcamInstance = new Webcam(camera, 'user', canvas)
}

async function start_camera() {
    if (!webcamInstance) initWebcam()
    
    // Initialize face detector if not ready
    if (!faceDetectorReady) {
        // Wait for module to load
        let attempts = 0;
        while (!window.FaceDetection && attempts < 50) {
            await new Promise(r => setTimeout(r, 100));
            attempts++;
        }
        
        if (window.FaceDetection) {
            faceDetectorReady = await window.FaceDetection.init()
        } else {
            console.warn("FaceDetection module not loaded")
        }
    }
    
    webcamInstance.start()
        .then(result => {
            Toastify({
                text: "Kamera berhasil diakses",
                duration: 2000,
                close: false,
                gravity: "top",
                position: "center",
                backgroundColor: "#4fbe87",
            }).showToast()

            $("#btn-open-camera").prop("disabled", false)
            
            // Setup face detection overlay and start loop
            var video = document.getElementById("camera")
            if (faceDetectorReady && video && window.FaceDetection) {
                window.FaceDetection.setupOverlay(video)
                window.FaceDetection.startLoop(video)
            }
        })
        .catch(err => {
            Toastify({
                text: "Terjadi kesalahan saat mengakses kamera",
                duration: 2000,
                close: false,
                gravity: "top",
                position: "center",
                backgroundColor: "#dc3545",
            }).showToast()
        })
}

function stop_camera() {
    if (window.FaceDetection) {
        window.FaceDetection.stopLoop()
    }
    
    if (webcamInstance) {
        webcamInstance.stop()
        webcamInstance = null
    }
}

// ============================================
// Modal Events for Webcam
// ============================================
function initWebcamModal() {
    // Start camera when modal opens
    $('#border-less').on('show.bs.modal', function () {
        start_camera()
    })

    // Stop camera when modal closes
    $('#border-less').on('hidden.bs.modal', function () {
        stop_camera()
    })
}

function take_snapshot() {
    // Validate face count before taking photo
    if (window.FaceDetection) {
        var validation = window.FaceDetection.validate()
        if (!validation.valid) {
            Toastify({
                text: validation.message,
                duration: 3000,
                close: false,
                gravity: "top",
                position: "center",
                backgroundColor: "#dc3545",
            }).showToast()
            return
        }
    }
    
    var video = document.getElementById("camera")
    var canvas = document.getElementById("result")
    var ctx = canvas.getContext("2d")
    
    // Get video dimensions
    var videoWidth = video.videoWidth
    var videoHeight = video.videoHeight
    
    // Calculate center crop 1:1
    var size = Math.min(videoWidth, videoHeight)
    var centerX = videoWidth / 2
    var centerY = videoHeight / 2
    var startX = centerX - size / 2
    var startY = centerY - size / 2
    
    // Set canvas size to square
    canvas.width = size
    canvas.height = size
    
    // Flip horizontal (mirror)
    ctx.translate(size, 0)
    ctx.scale(-1, 1)
    
    // Draw center cropped image (clean, no bounding boxes)
    ctx.drawImage(video, startX, startY, size, size, 0, 0, size, size)
    
    // Reset transform
    ctx.setTransform(1, 0, 0, 1, 0, 0)
    
    // Convert to base64
    var picture = canvas.toDataURL("image/png")
    $("#image").val(picture)

    $("#result").removeClass("d-none")
    $("#result-placeholder").addClass("d-none")

    Toastify({
        text: "Gambar berhasil diambil",
        duration: 2000,
        close: false,
        gravity: "top",
        position: "center",
        backgroundColor: "#4fbe87",
    }).showToast()
}

function restart_camera() {
    stop_camera()
    $("#result").addClass("d-none")
    $("#result-placeholder").removeClass("d-none")
    start_camera()
}

function usePhoto() {
    var picture = $("#image").val()
    if (picture) {
        $("#preview-img").attr("src", picture)
        $("#btn-open-camera").addClass("d-none")
        $("#camera-preview").removeClass("d-none")
        stop_camera()
    }
}

function resetPhoto() {
    $("#image").val("")
    $("#preview-img").attr("src", "")
    $("#btn-open-camera").removeClass("d-none")
    $("#camera-preview").addClass("d-none")
    $("#result").addClass("d-none")
    $("#result-placeholder").removeClass("d-none")
    var modal = new bootstrap.Modal(document.getElementById('border-less'))
    modal.show()
}

function resetForm() {
    $("#form-registrasi")[0].reset()
    $(".form-control, .form-select").removeClass("is-invalid is-valid")
    $(".invalid-feedback").remove()
    $("#image").val("")
    $("#preview-img").attr("src", "")
    $("#btn-open-camera").removeClass("d-none")
    $("#camera-preview").addClass("d-none")
    $("#result").addClass("d-none")
    $("#result-placeholder").removeClass("d-none")
    $("#row-pokja").addClass("hidden")
    $("#select-pokja").prop("required", false)
}

// ============================================
// Select Tujuan Change Handler
// ============================================
function initTujuanHandler() {
    $("#select-tujuan").change(function (e) {
        var value = e.currentTarget.value.toUpperCase()
        
        // Show Pokja dropdown when Tujuan is POKJA
        if (value === "POKJA") {
            $("#row-pokja").removeClass("hidden")
            $("#select-pokja").prop("required", true)
        } else {
            $("#row-pokja").addClass("hidden")
            $("#select-pokja").prop("required", false)
            $("#select-pokja").val("")
        }
    })
}

// ============================================
// Form Validation
// ============================================
function validateForm() {
    // Clear previous validation states
    $(".form-control, .form-select").removeClass("is-invalid is-valid")
    $(".invalid-feedback").remove()

    var fields = [
        { name: "nama", label: "Nama", type: "text", element: "input[name='nama']" },
        { name: "telepon", label: "No. Telepon", type: "phone", element: "input[name='telepon']" },
        { name: "instansi", label: "Instansi", type: "select", element: "#select-instansi" },
        { name: "skpd_opd", label: "SKPD/OPD", type: "text", element: "input[name='skpd_opd']" },
        { name: "tujuan", label: "Tujuan", type: "select", element: "#select-tujuan" },
    ]

    // Add Pokja if visible
    if (!$("#row-pokja").hasClass("hidden")) {
        fields.push({ name: "pokja", label: "Pokja", type: "select", element: "#select-pokja" })
    }

    // Add webcam validation
    fields.push({ name: "image", label: "Foto Webcam", type: "webcam", element: "#image" })

    // Validate sequentially
    for (var i = 0; i < fields.length; i++) {
        var field = fields[i]
        var $el = $(field.element)
        var value = $el.val()

        var isValid = false
        if (field.type === "phone") {
            isValid = value && value.length > 8
        } else if (field.type === "webcam") {
            isValid = value && value.length > 0
        } else {
            isValid = value && value.trim() !== ""
        }

        if (!isValid) {
            $el.addClass("is-invalid")
            $el.after('<div class="invalid-feedback">' + field.label + ' harus diisi</div>')
            
            // Scroll to first error
            $("html, body").animate({
                scrollTop: $el.offset().top - 100
            }, 300)

            // Show toast error
            Toastify({
                text: field.label + " harus diisi",
                duration: 2000,
                close: false,
                gravity: "top",
                position: "center",
                backgroundColor: "#dc3545",
            }).showToast()

            return false
        } else {
            $el.addClass("is-valid")
        }
    }

    return true
}

function validateFormPenyedia() {
    // Clear previous validation states
    $(".form-control, .form-select").removeClass("is-invalid is-valid")
    $(".invalid-feedback").remove()

    var fields = [
        { name: "nama", label: "Nama", type: "text", element: "input[name='nama']" },
        { name: "telepon", label: "No. Telepon", type: "phone", element: "input[name='telepon']" },
        { name: "company", label: "Perusahaan", type: "text", element: "input[name='company']" },
        { name: "tujuan", label: "Tujuan", type: "select", element: "#select-tujuan" },
    ]

    // Add Pokja if visible
    if (!$("#row-pokja").hasClass("hidden")) {
        fields.push({ name: "pokja", label: "Pokja", type: "select", element: "#select-pokja" })
    }

    // Add webcam validation
    fields.push({ name: "image", label: "Foto Webcam", type: "webcam", element: "#image" })

    // Validate sequentially
    for (var i = 0; i < fields.length; i++) {
        var field = fields[i]
        var $el = $(field.element)
        var value = $el.val()

        var isValid = false
        if (field.type === "phone") {
            isValid = value && value.length > 8
        } else if (field.type === "webcam") {
            isValid = value && value.length > 0
        } else {
            isValid = value && value.trim() !== ""
        }

        if (!isValid) {
            $el.addClass("is-invalid")
            $el.after('<div class="invalid-feedback">' + field.label + ' harus diisi</div>')
            
            // Scroll to first error
            $("html, body").animate({
                scrollTop: $el.offset().top - 100
            }, 300)

            // Show toast error
            Toastify({
                text: field.label + " harus diisi",
                duration: 2000,
                close: false,
                gravity: "top",
                position: "center",
                backgroundColor: "#dc3545",
            }).showToast()

            return false
        } else {
            $el.addClass("is-valid")
        }
    }

    return true
}

// ============================================
// Initialize on Document Ready
// ============================================
$(document).ready(function() {
    // Initialize phone numeric input
    $("input[name='telepon']").numeric({ decimal: false, negative: false })
    
    // Initialize Tujuan handler
    initTujuanHandler()
})
