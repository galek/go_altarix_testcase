задание 1. Написать демона для разбора очереди.

Начальные условия: 

●	Очередь сообщений определенной структуры (rabbitmq)

●	Таблица в базе (postgresql)


Структура сообщений:

●	access_token: required (string - pattern: ^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-(8|9|a|b)[a-f0-9]{3}-[a-f0-9]{12}$)

●	event_code: required (string - minLength: 1 - maxLength: 255)

●	stream_type: required (one of email, sms, push)

●	data: required (object) Произвольный набор полей

Необходимо написать демона для извлечения сообщений из очереди, их обработки и записи в базу. В зависимости от значения stream_type извлечь из объекта data атрибут по ключу person_* (* заменяет значение stream_type) и поместить его на верхний уровень сообщения с ключем to. Обработанное сообщение необходимо записать в базу с соответствующей структурой таблицы.


Пример исходного сообщения:
```
{
  "access_token": "0d10566b-7e7f-4c17-b2ea-f0e42a4df3c0",
  "event_code": "ispp",
  "stream_type": "email",
  "data": {
    "person_name": "Иван",
    "date": "2016-03-03",
    "person_email": "ivanivanov@gmail.com"
  }
}
```

Пример обработанного сообщения:
```
{
  "access_token": "0d10566b-7e7f-4c17-b2ea-f0e42a4df3c0",
  "event_code": "ispp",
  "stream_type": "email",
  "to": "ivanivanov@gmail.com"
  "data": {
    "person_name": "Иван",   
    "date": "2016-03-03"
  }
}
```

Важные моменты:

●	Проследить за корректной работой скрипта в аварийных ситуациях

●	(Дополнительно) Добавить Dockerfile



