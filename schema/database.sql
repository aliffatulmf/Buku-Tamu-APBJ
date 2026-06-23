-- Buku Tamu APBJ Database Schema
-- MySQL

CREATE TABLE IF NOT EXISTS destinations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    destination_name VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS consultations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    consultation_name VARCHAR(255) NOT NULL UNIQUE,
    destination_id INT DEFAULT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_destination_id (destination_id),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (destination_id) REFERENCES destinations(id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS pokjas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pokja_name VARCHAR(255) NOT NULL UNIQUE,
    status TINYINT(1) DEFAULT 1,
    destination_id INT DEFAULT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_destination_id (destination_id),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (destination_id) REFERENCES destinations(id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS agencies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    agency_name VARCHAR(255) NOT NULL UNIQUE,
    agency_email VARCHAR(255) DEFAULT NULL,
    agency_telephone VARCHAR(50) DEFAULT NULL,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS pemdas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pemda_name VARCHAR(255) NOT NULL,
    phone VARCHAR(50) DEFAULT NULL,
    skpd_opd VARCHAR(255) NOT NULL,
    agency_id INT NOT NULL,
    destination VARCHAR(255) DEFAULT NULL,
    consultation VARCHAR(255) DEFAULT NULL,
    pokja VARCHAR(255) DEFAULT NULL,
    image VARCHAR(255) NOT NULL,
    verified TINYINT(1) NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_agency_id (agency_id),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (agency_id) REFERENCES agencies(id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS providers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    provider_name VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(50) DEFAULT NULL,
    company VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    destination VARCHAR(255) NOT NULL,
    consultation VARCHAR(255) DEFAULT NULL,
    pokja VARCHAR(255) DEFAULT NULL,
    image VARCHAR(255) NOT NULL,
    verified TINYINT(1) NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
