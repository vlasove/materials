# Шажки

***Ссылка:*** https://losst.ru/ustanovka-postgresql-ubuntu-16-04

## Шаг 1. Установка
* ```sudo apt-get update```
* ```sudo apt install postgresql postgresql-contrib```

## Шаг 2. Настройка
* ```sudo -i -u postgres``` - открытие терминала и переключение на пользователя ```postgres``` (```exit``` для завершения) 
* Дефолтная учетка
* ```psql``` - консоль управления (```\q``` выход)
    * ```\conninfo``` - конфиги подключения
    * ```\q``` - выйти
* Создание роли ```postgres```
    * ```createuser --interactive```

#### Важно
* Точно также как имена ролей сопоставляются с системными пользователями, имя базы данных будет подбираться по имени пользователя

* ```createdb <role_name>```
* Подкючимся к бд - войти от имени ```<role_name>``` : ```sudo su - <role_name>```
* Подключение к другой бд ```psql -d postgres```

### Шаг 3. Операции. Простейшие
* ```createdb tester``` - создать бд ```tester```
* ```psql -d tesster``` - подключение к бд ```tester```
* Запрос на создание таблицы 
```
CREATE TABLE playground (
equip_id serial PRIMARY KEY,
type varchar (50) NOT NULL,
color varchar (25) NOT NULL,
location varchar(25) check (location in ('north', 'south', 'west', 'east', 'northeast', 'southeast', 'southwest', 'northwest')),
install_date date
);
```
* ```\d``` - вывести все что есть в бд
    * ```\dt``` - вывести только таблицы
    * ```\ds``` - вывести все сиквенс (переменные)
* Вставка данных
```
INSERT INTO playground (type, color, location, install_date) VALUES ('slide', 'blue', 'south', '2016-04-28');
INSERT INTO playground (type, color, location, install_date) VALUES ('swing', 'yellow', 'northwest', '2015-08-16');
```
* Чтение ```SELECT * FROM playground;```
* Удаление ```DELETE FROM playground WHERE type='slide';```