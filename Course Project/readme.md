## Проект

телеграм бот для улучшения голосовух

## Фичи

### Основные

* нормализация громкости
* выбор языка интерфейса
* сохранение порядка сообщений
* ответ по мере готовности или батча целиком? - по мере готовности. Не удобно детектить батч.


### Экстра

* rate limiting
* работа без использования subprocess и дампов в файлы
* вести статистику по юзерам, запросам, объёмам
* доп усложнение: биллинг, регистрация, профиль, настройки
* выравнивание громкости (компрессия)
* распознование речи
* выбор языка распознования речи
* прогон через нейросетевой улучшайзер
* реклама
* контекстная реклама

## флоу

Онбординг:

* юзер жмёт старт
* бот отвечает приветствием, линк/кнопка смены языка UI
* создаём профиль юзера

Флоу:

* юзер присылает/форвардит голосовуху/видеокружок
* бот читает профиль юзера, настройки ui, методы обработки
* бот процессит сообщение и возвращает результат
* в том же порядке в котором получил сообщения

Платные фичи:
* юзер выбирает платную фичу
* присылаем линк на оплату
* сохраняем оплаченный заказ

### схема бд

User (tg-id)
profile (lang, is_recognize_speech, is_ML_enchancer, is_compressor)

### технические решения
Нормализация громоксти через ffmpeg

Профиль юзера с настройками в БД

Язык: go/ts/c#?
