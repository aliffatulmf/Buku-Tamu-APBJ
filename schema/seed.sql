-- Seed Data for Buku Tamu APBJ
-- SQLite

-- Destinations (Tujuan)
INSERT INTO destinations (id, destination_name) VALUES
('advokasi', 'Advokasi'),
('lpse', 'LPSE'),
('pokja', 'POKJA');

-- Consultations (Konsultasi)
INSERT INTO consultations (id, consultation_name, destination_id) VALUES
('konsultasi_spse', 'Konsultasi SPSE', 'lpse'),
('konsultasi_sirup', 'Konsultasi Sirup', 'lpse'),
('konsultasi_blangkon_jateng', 'Konsultasi Blangkon Jateng', 'lpse');

-- Pokjas
INSERT INTO pokjas (id, pokja_name, status, destination_id) VALUES
('pokja_pemilihan_1a', 'Pokja Pemilihan 1a', 1, 'pokja'),
('pokja_pemilihan_2a', 'Pokja Pemilihan 2a', 1, 'pokja'),
('pokja_pemilihan_3a', 'Pokja Pemilihan 3a', 1, 'pokja'),
('pokja_pemilihan_4a', 'Pokja Pemilihan 4a', 1, 'pokja'),
('pokja_pemilihan_5a', 'Pokja Pemilihan 5a', 1, 'pokja'),
('pokja_pemilihan_6a', 'Pokja Pemilihan 6a', 1, 'pokja'),
('pokja_pemilihan_7a', 'Pokja Pemilihan 7a', 1, 'pokja'),
('pokja_pemilihan_8a', 'Pokja Pemilihan 8a', 1, 'pokja'),
('pokja_pemilihan_9a', 'Pokja Pemilihan 9a', 1, 'pokja'),
('pokja_pemilihan_10a', 'Pokja Pemilihan 10a', 1, 'pokja'),
('pokja_pemilihan_11a', 'Pokja Pemilihan 11a', 1, 'pokja'),
('pokja_pemilihan_12a', 'Pokja Pemilihan 12a', 1, 'pokja'),
('pokja_pemilihan_13a', 'Pokja Pemilihan 13a', 1, 'pokja');

-- Agencies (Instansi/OPD)
INSERT INTO agencies (agency_name, agency_email, agency_telephone) VALUES
('Pemerintah Provinsi Jawa Tengah', NULL, NULL),
('Pemerintah Kabupaten Cilacap', NULL, NULL),
('Pemerintah Kabupaten Banyumas', NULL, NULL),
('Pemerintah Kabupaten Purbalingga', NULL, NULL),
('Pemerintah Kabupaten Banjarnegara', NULL, NULL),
('Pemerintah Kabupaten Kebumen', NULL, NULL),
('Pemerintah Kabupaten Purworejo', NULL, NULL),
('Pemerintah Kabupaten Wonosobo', NULL, NULL),
('Pemerintah Kabupaten Magelang', NULL, NULL),
('Pemerintah Kabupaten Boyolali', NULL, NULL),
('Pemerintah Kabupaten Klaten', NULL, NULL),
('Pemerintah Kabupaten Sukoharjo', NULL, NULL),
('Pemerintah Kabupaten Wonogiri', NULL, NULL),
('Pemerintah Kabupaten Karanganyar', NULL, NULL),
('Pemerintah Kabupaten Seragen', NULL, NULL),
('Pemerintah Kabupaten Grobogan', NULL, NULL),
('Pemerintah Kabupaten Blora', NULL, NULL),
('Pemerintah Kabupaten Rembang', NULL, NULL),
('Pemerintah Kabupaten Pati', NULL, NULL),
('Pemerintah Kabupaten Kudus', NULL, NULL),
('Pemerintah Kabupaten Jepara', NULL, NULL),
('Pemerintah Kabupaten Demak', NULL, NULL),
('Pemerintah Kabupaten Semarang', NULL, NULL),
('Pemerintah Kabupaten Temanggung', NULL, NULL),
('Pemerintah Kabupaten Kendal', NULL, NULL),
('Pemerintah Kabupaten Batang', NULL, NULL),
('Pemerintah Kabupaten Pekalongan', NULL, NULL),
('Pemerintah Kabupaten Pemalang', NULL, NULL),
('Pemerintah Kabupaten Tegal', NULL, NULL),
('Pemerintah Kabupaten Brebes', NULL, NULL),
('Pemerintah Kota Magelang', NULL, NULL),
('Pemerintah Kota Surakarta', NULL, NULL),
('Pemerintah Kota Salatiga', NULL, NULL),
('Pemerintah Kota Semarang', NULL, NULL),
('Pemerintah Kota Pekalongan', NULL, NULL),
('Pemerintah Kota Tegal', NULL, NULL);

-- Sample Pemda Data (for testing)
INSERT INTO pemdas (pemda_name, phone, skpd_opd, agency_id, destination, consultation, pokja, image, verified) VALUES
('Budi Santoso', '081234567890', 'Dinas Pendidikan', 1, 'advokasi', NULL, NULL, 'media/img/test-pemda-1.png', 1),
('Siti Rahayu', '082345678901', 'Dinas Kesehatan', 2, 'lpse', 'konsultasi_spse', NULL, 'media/img/test-pemda-2.png', 0),
('Ahmad Hidayat', '083456789012', 'Dinas PU', 3, 'pokja', NULL, 'pokja_pemilihan_1a', 'media/img/test-pemda-3.png', 1);

-- Sample Provider Data (for testing)
INSERT INTO providers (provider_name, phone, company, description, destination, consultation, pokja, image, verified) VALUES
('PT Teknologi Nusantara', '085678901234', 'PT Teknologi Nusantara', 'Perusahaan software', 'advokasi', NULL, NULL, 'media/img/test-provider-1.png', 1),
('CV Digital Solusi', '086789012345', 'CV Digital Solusi', 'Konsultan IT', 'lpse', 'konsultasi_sirup', NULL, 'media/img/test-provider-2.png', 0),
('PT Konstruksi Jaya', '087890123456', 'PT Konstruksi Jaya', 'Kontraktor', 'pokja', NULL, 'pokja_pemilihan_2a', 'media/img/test-provider-3.png', 1);
