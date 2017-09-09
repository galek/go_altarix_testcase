/*
Таблица которая задает уникальные имена
Формат:
 uid(autoincrement)-Name

Создается для схемы public
*/


CREATE TABLE public.ID_NAMES(  
   uid serial NOT NULL PRIMARY KEY,
   name TEXT
);

INSERT INTO public.ID_NAMES (name) VALUES  
('Иван'),
('Николай'),
('Андрей'),
('Алексей'),
('Саша'),
('Даша')
;

/*
Таблица уникальных кодов
Формат: 
Код-Описание
*/
CREATE TABLE public.EVENT_CODES(  
   uid serial NOT NULL PRIMARY KEY,
   descr TEXT
);

INSERT INTO public.EVENT_CODES (descr) VALUES  
('Загрузка'),
('Вышло новое обновление'),
('Ваша очередь'),
('Клиническая смерть сервера'),
('Клиническая смерть клиента'),
('Слушает музыку')
;

/*
Таблица типов stream
Формат: 
Код-Описание
*/
CREATE TABLE public.STREAM_TYPES(  
   uid serial NOT NULL PRIMARY KEY,
   descr TEXT
);

INSERT INTO public.STREAM_TYPES (descr) VALUES  
('SMS'),
('EMAIL'),
('PUSH')
;

/*
Таблица uud_email
Формат: 
Код-почта
*/
CREATE TABLE public.UUID_EMAIL(  
   uid serial NOT NULL PRIMARY KEY,
   email TEXT
);

INSERT INTO public.UUID_EMAIL (email) VALUES  
('test@test.ru'),
('test@test2.ru'),
('test@test3.ru'),
('test@test4.ru'),
('test@test5.ru')
;


/*
Таблица uud_sms
uud_push
*/

/*таблица access_token токенов*/

/*
Результирующая таблица
*/
CREATE TABLE public.ResultTable(  
access_token int,
event_token int, 
stream_type int,
/*Data block*/
person_name int,
person_email int,
FOREIGN KEY (event_token) REFERENCES public.EVENT_CODES(uid),
FOREIGN KEY (stream_type) REFERENCES public.STREAM_TYPES(uid),
FOREIGN KEY (person_email) REFERENCES public.UUID_EMAIL(uid),
FOREIGN KEY (person_name) REFERENCES public.ID_NAMES(uid),
/*обычная дата DATE*/
person_date date
);