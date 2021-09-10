## JSON

Научиться работать с ```JSON```:
    * [x] Научиться marshall\unmarshall из байтов
    * [x] Научиться заполнять структуры , научиться из структур делать байтовые json
    * [x] Понять, какие фичи есть у тегов json
    * [ ] Научиться работать с файлами ```.json```
    * [ ] Научиться валидировать ```json```
    * [ ] Написать небольшое приложение, которое обойдет директорию, найдет все ```json``` и соберет их в один файл
    * [ ] Написать тесты к этому приложению
    * [ ] Добавить graceful shutdown приложению
    * [ ] Написать тесты

### Небольшое приложение
Хочу содать небольшое приложение, которое будет обладать седующим функционалом:
* ```./app --validate=path/to/data.json --schema=path/to/schema.json``` - проверяет, что ```data.json``` удовлетворяет схеме ```schema.json```

* ```./app --validate=path/to/dir --schema=path/to/schema.json``` - возвращает набор файлов в директории, которые удовлетворяеют схеме
* ```./app --copy=path/to/dir --validate=path/to/data.json --schema=path/to/schema.json``` - переносит файл data.json в директорию ```path/to/dir``` , если он удовлетворяет схеме  ```schema.json```

* ```./app --copy=path/to/dir --vaidate=path/to/dir2 --schema=path/to/schema.json``` - переносит все файлы из директории ```path/to/dir2``` в директорию ```copy```, если они удовлетворяют схеме
* ```./app --output=path/to/output.json --input=path/to/dir schema=path/to/schema.json``` - собирает все json файлы из директории ```input``` в один большой ```output``` если они удовлетворяют схеме ```schema.json```.

#### Аргументы CLI
