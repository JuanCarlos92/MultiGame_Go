-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 10-06-2025 a las 13:08:33
-- Versión del servidor: 10.4.32-MariaDB
-- Versión de PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `juego`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `jugador`
--

CREATE TABLE `jugador` (
  `id` int(11) NOT NULL,
  `Name` longtext DEFAULT NULL,
  `Email` longtext DEFAULT NULL,
  `Password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `jugador`
--

INSERT INTO `jugador` (`id`, `Name`, `Email`, `Password`) VALUES
(1, 'ejemplo1 ', 'ejemplo1@gmail.com', '$2a$10$KKYo0tKZXch8pdhFjjFapurzwNirmngSZ9C642WtlYJynGi5QHHxy'),
(2, 'ejemplo', 'ejemplo@gmail.com', '$2a$10$axhPaRzUD7leBMENXXY8HOo3WF0AcXCbbMyDLfYqvWWg./zIALgWe'),
(3, 'ejemplo2', 'ejemplo2@gmail.com', '$2a$10$Y4w92H0AOtf4C816tvV01.ST/jWC8PV/.MhOs.ODEytc6VlmEG4K6'),
(4, 'pruebaingles', 'ingles@gmail.com', '$2a$10$n2X0ceCBRGx0I7VLIVTEp.v7rrRrQaJCLbJTpdRAw0qVNGgQU4O2.');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `jugador`
--
ALTER TABLE `jugador`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `jugador`
--
ALTER TABLE `jugador`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
