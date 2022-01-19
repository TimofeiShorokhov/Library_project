-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Янв 19 2022 г., 17:04
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

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
(12, 'Третьяков', 'authors/Третьяков.jpg'),
(14, 'Ершов', 'authors/Ершов.jpg');

-- --------------------------------------------------------

--
-- Структура таблицы `books`
--

CREATE TABLE `books` (
  `book_id` int UNSIGNED NOT NULL,
  `book_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `year` int NOT NULL,
  `quantity` int NOT NULL,
  `available` int NOT NULL,
  `registration` date NOT NULL DEFAULT '1999-12-31',
  `book_price` int NOT NULL,
  `Image_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `books`
--

INSERT INTO `books` (`book_id`, `book_name`, `year`, `quantity`, `available`, `registration`, `book_price`, `Image_path`) VALUES
(1, 'Колец', 1996, 5, 5, '2022-01-10', 350, 'img/Колец.jpg'),
(2, 'Оно', 1989, 5, 5, '2022-01-11', 450, 'img/Оно.jpg'),
(3, 'Гарри Поттер', 1989, 5, 5, '2022-01-11', 450, 'img/Гарри_Поттер.jpg'),
(4, 'Тестовая книга', 2011, 5, 5, '2022-01-19', 250, 'img/Тестовая_книга.jpg'),
(5, 'Тест обложки', 2011, 5, 5, '2022-01-19', 250, 'img/Тест_обложки.jpg');

-- --------------------------------------------------------

--
-- Структура таблицы `book_authors`
--

CREATE TABLE `book_authors` (
  `book_id` int UNSIGNED NOT NULL,
  `author_id` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `book_authors`
--

INSERT INTO `book_authors` (`book_id`, `author_id`) VALUES
(1, 1),
(4, 1),
(5, 1),
(1, 2),
(3, 2),
(4, 2),
(5, 2),
(1, 3),
(2, 3),
(4, 3),
(5, 3),
(1, 4),
(4, 4),
(5, 4),
(4, 5),
(5, 5),
(4, 6),
(5, 6);

-- --------------------------------------------------------

--
-- Структура таблицы `book_genre`
--

CREATE TABLE `book_genre` (
  `book_id` int UNSIGNED NOT NULL,
  `genre_id` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `book_genre`
--

INSERT INTO `book_genre` (`book_id`, `genre_id`) VALUES
(1, 1),
(1, 2),
(2, 4),
(2, 5),
(3, 2),
(3, 3),
(3, 5),
(3, 6),
(4, 1),
(4, 2),
(4, 3),
(4, 4),
(4, 5),
(4, 6),
(5, 1),
(5, 2),
(5, 3),
(5, 4),
(5, 5),
(5, 6);

-- --------------------------------------------------------

--
-- Структура таблицы `documents`
--

CREATE TABLE `documents` (
  `id` int UNSIGNED NOT NULL,
  `reader_surname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `book_id` int NOT NULL,
  `book_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
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
(1, 'Комедия'),
(2, 'Приключения'),
(3, 'Боевик'),
(4, 'Ужасы'),
(5, 'Драма'),
(6, 'Фантастика');

-- --------------------------------------------------------

--
-- Структура таблицы `instances`
--

CREATE TABLE `instances` (
  `instance_id` int UNSIGNED NOT NULL,
  `instance_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `damage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'Нет повреждений',
  `instance_price` int NOT NULL,
  `return_date` date NOT NULL DEFAULT '1999-12-31',
  `dmg_photo` varchar(255) NOT NULL DEFAULT 'Книгу еще не брали',
  `rating` int NOT NULL DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `instances`
--

INSERT INTO `instances` (`instance_id`, `instance_name`, `damage`, `instance_price`, `return_date`, `dmg_photo`, `rating`) VALUES
(295, 'Тест обложки', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(292, 'Тест обложки', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(294, 'Тест обложки', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(293, 'Тест обложки', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(280, 'Тестовая книга', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(291, 'Тест обложки', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(279, 'Тестовая книга', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(278, 'Тестовая книга', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(277, 'Тестовая книга', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(276, 'Тестовая книга', 'Нет повреждений', 250, '1999-12-31', 'Книгу еще не брали', 0),
(275, 'Гарри Поттер', 'Помята обложка', 450, '2022-01-19', 'ссылка на фото', 3),
(274, 'Гарри Поттер', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(273, 'Гарри Поттер', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(272, 'Гарри Поттер', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(271, 'Гарри Поттер', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(190, 'Оно', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(189, 'Оно', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(188, 'Оно', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(187, 'Оно', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(186, 'Оно', 'Нет повреждений', 450, '1999-12-31', 'Книгу еще не брали', 0),
(185, 'Колец', 'Помята обложка', 450, '2022-01-10', 'ссылка на фото', 3),
(184, 'Колец', 'Нет повреждений', 350, '1999-12-31', 'Книгу еще не брали', 0),
(183, 'Колец', 'Нет повреждений', 350, '1999-12-31', 'Книгу еще не брали', 0),
(182, 'Колец', 'Нет повреждений', 350, '1999-12-31', 'Книгу еще не брали', 0),
(181, 'Колец', 'Нет повреждений', 350, '1999-12-31', 'Книгу еще не брали', 0);

-- --------------------------------------------------------

--
-- Структура таблицы `readers`
--

CREATE TABLE `readers` (
  `id` int UNSIGNED NOT NULL,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `birthdate` date DEFAULT NULL,
  `adress` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `surname` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `debt` int DEFAULT '0'
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `readers`
--

INSERT INTO `readers` (`id`, `name`, `birthdate`, `adress`, `surname`, `email`, `debt`) VALUES
(1, 'Леонид', '1998-02-13', 'улица Советская', 'Балычев', 'leno@yandex.com', 0),
(2, 'Евгений', '1987-02-02', 'улица Богдановича', 'Михайлов', 'evg@mail.ru', 0),
(3, 'Тимофей', '1997-12-17', 'ул Богдановича 78, 439', 'Шорохов', 'animaster-animaster@mail.ru', 0),
(4, 'Ксения', '2002-10-01', 'ул Макаенка 12В', 'Лесникова', 'lesnikova@gmail.com', 0),
(5, 'Егор', '1997-07-23', 'минск', 'Янкин', 'yank97@mail.ru', 0),
(9, 'Владимир', '1967-09-05', 'ул. Мястровская 15А', 'Бодров', 'bodrov@mail.ru', 0),
(35, 'Marcus', '2023-11-13', 'London, UK, Brighton street 53', 'Rashford', 'rashford@gmail.com', 0);

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
-- Индексы таблицы `book_authors`
--
ALTER TABLE `book_authors`
  ADD PRIMARY KEY (`book_id`,`author_id`),
  ADD KEY `author_id` (`author_id`);

--
-- Индексы таблицы `book_genre`
--
ALTER TABLE `book_genre`
  ADD PRIMARY KEY (`book_id`,`genre_id`),
  ADD UNIQUE KEY `genre_id` (`genre_id`,`book_id`);

--
-- Индексы таблицы `documents`
--
ALTER TABLE `documents`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `genres`
--
ALTER TABLE `genres`
  ADD PRIMARY KEY (`genre_id`);

--
-- Индексы таблицы `instances`
--
ALTER TABLE `instances`
  ADD PRIMARY KEY (`instance_id`);

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
  MODIFY `author_id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT для таблицы `documents`
--
ALTER TABLE `documents`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=341;

--
-- AUTO_INCREMENT для таблицы `instances`
--
ALTER TABLE `instances`
  MODIFY `instance_id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=296;

--
-- AUTO_INCREMENT для таблицы `readers`
--
ALTER TABLE `readers`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `book_authors`
--
ALTER TABLE `book_authors`
  ADD CONSTRAINT `book_authors_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  ADD CONSTRAINT `book_authors_ibfk_2` FOREIGN KEY (`author_id`) REFERENCES `authors` (`author_id`) ON DELETE CASCADE ON UPDATE RESTRICT;

--
-- Ограничения внешнего ключа таблицы `book_genre`
--
ALTER TABLE `book_genre`
  ADD CONSTRAINT `book_genre_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`genre_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `book_genre_ibfk_3` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`) ON DELETE CASCADE ON UPDATE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
