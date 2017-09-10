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
Таблица uud_email
Формат: 
Код-номер телефона
*/
CREATE TABLE public.UUID_SMS(  
   uid serial NOT NULL PRIMARY KEY,
   tel_number TEXT
);

INSERT INTO public.UUID_SMS (tel_number) VALUES  
('000000001'),
('000000002'),
('000000003'),
('000000004'),
('000000005')
;

/*
Таблица uud_push
Формат: 
Код-заголовок уведомления, подробный текст, иконка
*/
CREATE TABLE public.UUID_PUSH(  
   uid serial NOT NULL PRIMARY KEY,
   caption TEXT,
   body TEXT,
   icon int
);

INSERT INTO public.UUID_PUSH (caption, body, icon) VALUES  
('cap1','body1', 0),
('cap2','body2', 0),
('cap3','body3', 0),
('cap4','body4', 0)
;


/*
таблица access_token токенов
SELECT md5(random()::text || clock_timestamp()::text)::uuid
на выходе дает 
38b2cfb8-eb40-fc3d-9a81-49304b21cdb6

на моей psql уже готовые функции типа uuid_generate_v4, не работают :(
*/
CREATE TABLE public.ACCESS_TOKENS(  
   uid serial NOT NULL PRIMARY KEY,
   TOKEN TEXT
);

INSERT INTO public.ACCESS_TOKENS (TOKEN) VALUES  
('38b2cfb8-eb40-fc3d-9a81-49304b21cdb6')
;

/*
Результирующая таблица
*/
CREATE TABLE public.ResultTable(  
access_token int NOT NULL,
event_token int NOT NULL, 
stream_type int NOT NULL,
/*Data block*/
person_name int NOT NULL,
person_email int,
person_phone int,
FOREIGN KEY (access_token) REFERENCES public.ACCESS_TOKENS(uid),
FOREIGN KEY (event_token) REFERENCES public.EVENT_CODES(uid),
FOREIGN KEY (stream_type) REFERENCES public.STREAM_TYPES(uid),
FOREIGN KEY (person_email) REFERENCES public.UUID_EMAIL(uid),
FOREIGN KEY (person_phone) REFERENCES public.UUID_SMS(uid),
FOREIGN KEY (person_name) REFERENCES public.ID_NAMES(uid),
/*обычная дата DATE*/
person_date date
);
/*
Результирующая таблица
*/
CREATE TABLE public.ToTable(  
access_token int NOT NULL,
event_token int NOT NULL, 
stream_type int NOT NULL,
/*Data block*/
person_name int NOT NULL,
person_to TEXT,
FOREIGN KEY (access_token) REFERENCES public.ACCESS_TOKENS(uid),
FOREIGN KEY (event_token) REFERENCES public.EVENT_CODES(uid),
FOREIGN KEY (stream_type) REFERENCES public.STREAM_TYPES(uid),
FOREIGN KEY (person_name) REFERENCES public.ID_NAMES(uid),
/*обычная дата DATE*/
person_date date
);