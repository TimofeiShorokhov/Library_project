-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Янв 04 2022 г., 18:31
-- Версия сервера: 8.0.24
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `library`
--

-- --------------------------------------------------------

--
-- Структура таблицы `authors`
--

CREATE TABLE `authors` (
  `author_id` int UNSIGNED NOT NULL,
  `author_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author_image` varchar(200) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `authors`
--

INSERT INTO `authors` (`author_id`, `author_name`, `author_image`) VALUES
(1, 'Достоевский', '0'),
(2, 'Роулинг', '0'),
(3, 'Кинг', '0'),
(4, 'Уотсон', '0'),
(5, 'Авдеева', '0'),
(6, 'Толкин', '0'),
(7, 'Дойл', '0'),
(8, 'Ремарк', '0'),
(9, 'Лондон', '0'),
(10, 'Кафка', '0'),
(11, 'Пушкин', '0'),
(12, 'Третьяков', './authorImage/Третьяков.jpg');

-- --------------------------------------------------------

--
-- Структура таблицы `books`
--

CREATE TABLE `books` (
  `book_id` int UNSIGNED NOT NULL,
  `book_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `book_genre_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `book_author_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int NOT NULL,
  `quantity` int NOT NULL,
  `available` int NOT NULL DEFAULT '0',
  `registration` date NOT NULL DEFAULT '1999-12-31',
  `book_price` int NOT NULL,
  `Image_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `books`
--

INSERT INTO `books` (`book_id`, `book_name`, `book_genre_id`, `book_author_id`, `year`, `quantity`, `available`, `registration`, `book_price`, `Image_path`) VALUES
(56, 'Властелин Колец', '[Боевик Фантастика Приключения Сказка Фэнтези]', '[Толкин]', 1954, 15, 15, '2021-12-27', 500, './img/Властелин Колец.jpg'),
(53, 'Оно', '[Драма Ужасы Фэнтези]', '[Кинг]', 1995, 15, 15, '2021-12-25', 300, './img/Оно.jpg'),
(52, 'Гарри Поттер', '[Боевик Фантастика Приключения Сказка Фэнтези]', '[Роулинг]', 1997, 15, 15, '2021-12-25', 450, './img/Гарри Поттер.jpg'),
(51, 'Хоббит', '[Боевик Комедия Фантастика Приключения Сказка Фэнтези]', '[Толкин]', 1937, 15, 15, '2021-12-25', 550, './img/Хоббит.jpg'),
(116, 'Что-то', '[Боевик Фантастика Приключения Сказка Фэнтези]', 'Роулинг', 1996, 15, 15, '2022-01-03', 500, './images/book_img/Что-то.jpg'),
(122, 'Книга', 'Боевик, Комедия, Приключения, Сказка, ', 'Роулинг, Кинг, ', 1996, 15, 15, '2022-01-04', 500, './images/book_img/Книга.jpg');

-- --------------------------------------------------------

--
-- Структура таблицы `documents`
--

CREATE TABLE `documents` (
  `id` int UNSIGNED NOT NULL,
  `reader_surname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `book` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `date` date NOT NULL,
  `price` bigint NOT NULL,
  `quant` int NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

-- --------------------------------------------------------

--
-- Структура таблицы `genres`
--

CREATE TABLE `genres` (
  `genre_id` int UNSIGNED NOT NULL,
  `book_genre` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `genres`
--

INSERT INTO `genres` (`genre_id`, `book_genre`) VALUES
(1, 'Боевик'),
(2, 'Комедия'),
(3, 'Драма'),
(4, 'Фантастика'),
(5, 'Приключения'),
(6, 'Биография'),
(7, 'Сказка'),
(8, 'Ужасы'),
(10, 'Фэнтези'),
(9, 'Роман');

-- --------------------------------------------------------

--
-- Структура таблицы `readers`
--

CREATE TABLE `readers` (
  `id` int UNSIGNED NOT NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `birthdate` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `adress` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `surname` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `debt` int DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `readers`
--

INSERT INTO `readers` (`id`, `name`, `birthdate`, `adress`, `surname`, `email`, `debt`) VALUES
(1, 'Леонид', '1962-02-05', 'улица Советская', 'Балычев', 'leno@yandex.com', 0),
(2, 'Евгений', '1998-08-23', 'улица Богдановича', 'Михайлов', 'evg@mail.ru', 0),
(3, 'Тимофей', '17 декабря 1997 года', 'ул Богдановича 78, 439', 'Шорохов', 'limjkeee@mail.ru', 0),
(4, 'Ксения', 'какой-то день какой-то месяц 2002 года', 'ул Макаенка 12В', 'Лесникова', 'lesnikova@gmail.com', 0),
(5, 'Егор', '97 год', 'минск', 'Янкин', 'какой-то', 0),
(9, 'Владимир', '10.05.1985', 'ул. Мястровская 15А', 'Бодров', 'bodrov@mail.ru', 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `authors`
--
ALTER TABLE `authors`
  ADD PRIMARY KEY (`author_id`);

--
-- Индексы таблицы `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`book_id`);

--
-- Индексы таблицы `documents`
--
ALTER TABLE `documents`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `readers`
--
ALTER TABLE `readers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `authors`
--
ALTER TABLE `authors`
  MODIFY `author_id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT для таблицы `books`
--
ALTER TABLE `books`
  MODIFY `book_id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=123;

--
-- AUTO_INCREMENT для таблицы `documents`
--
ALTER TABLE `documents`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=227;

--
-- AUTO_INCREMENT для таблицы `readers`
--
ALTER TABLE `readers`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
