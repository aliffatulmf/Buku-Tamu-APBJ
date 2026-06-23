-- Seed Data for Buku Tamu APBJ
-- MySQL

-- Destinations (Tujuan)
INSERT INTO destinations (id, destination_name, created_at, updated_at) VALUES
(1, 'Advokasi', NOW(), NOW()),
(2, 'LPSE', NOW(), NOW()),
(3, 'POKJA', NOW(), NOW());

-- Consultations (Konsultasi)
INSERT INTO consultations (consultation_name, destination_id, created_at, updated_at) VALUES
('Konsultasi SPSE', 2, NOW(), NOW()),
('Konsultasi Sirup', 2, NOW(), NOW()),
('Konsultasi Blangkon Jateng', 2, NOW(), NOW());

-- Pokjas
INSERT INTO pokjas (pokja_name, status, destination_id, created_at, updated_at) VALUES
('Pokja Pemilihan 1a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 2a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 3a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 4a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 5a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 6a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 7a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 8a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 9a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 10a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 11a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 12a', 1, 3, NOW(), NOW()),
('Pokja Pemilihan 13a', 1, 3, NOW(), NOW());

-- Agencies (Instansi/OPD)
INSERT INTO agencies (agency_name, agency_email, agency_telephone, created_at, updated_at) VALUES
('Pemerintah Provinsi Jawa Tengah', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Cilacap', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Banyumas', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Purbalingga', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Banjarnegara', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Kebumen', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Purworejo', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Wonosobo', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Magelang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Boyolali', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Klaten', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Sukoharjo', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Wonogiri', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Karanganyar', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Seragen', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Grobogan', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Blora', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Rembang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Pati', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Kudus', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Jepara', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Demak', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Semarang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Temanggung', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Kendal', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Batang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Pekalongan', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Pemalang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Tegal', NULL, NULL, NOW(), NOW()),
('Pemerintah Kabupaten Brebes', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Magelang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Surakarta', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Salatiga', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Semarang', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Pekalongan', NULL, NULL, NOW(), NOW()),
('Pemerintah Kota Tegal', NULL, NULL, NOW(), NOW());

-- Sample Pemda Data (for testing)
INSERT INTO pemdas (pemda_name, phone, skpd_opd, agency_id, destination, consultation, pokja, image, verified, created_at, updated_at) VALUES
('Budi Santoso', '081234567890', 'Dinas Pendidikan', 1, 'Advokasi', NULL, NULL, 'media/img/test-pemda-1.png', 1, NOW(), NOW()),
('Siti Rahayu', '082345678901', 'Dinas Kesehatan', 2, 'LPSE', 'Konsultasi SPSE', NULL, 'media/img/test-pemda-2.png', 0, NOW(), NOW()),
('Ahmad Hidayat', '083456789012', 'Dinas PU', 3, 'POKJA', NULL, 'Pokja Pemilihan 1a', 'media/img/test-pemda-3.png', 1, NOW(), NOW());

-- Sample Provider Data (for testing)
INSERT INTO providers (provider_name, phone, company, description, destination, consultation, pokja, image, verified, created_at, updated_at) VALUES
('PT Teknologi Nusantara', '085678901234', 'PT Teknologi Nusantara', 'Perusahaan software', 'Advokasi', NULL, NULL, 'media/img/test-provider-1.png', 1, NOW(), NOW()),
('CV Digital Solusi', '086789012345', 'CV Digital Solusi', 'Konsultan IT', 'LPSE', 'Konsultasi Sirup', NULL, 'media/img/test-provider-2.png', 0, NOW(), NOW()),
('PT Konstruksi Jaya', '087890123456', 'PT Konstruksi Jaya', 'Kontraktor', 'POKJA', NULL, 'Pokja Pemilihan 2a', 'media/img/test-provider-3.png', 1, NOW(), NOW());
