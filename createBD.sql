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