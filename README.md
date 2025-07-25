# PersonalCard

Консольний веб-додаток **PersonalCard**, який дозволяє користувачу створювати особисту картотеку з предметами, оцінками та нотатками.

## Функціонал

Програма має простий та зручний функціонал:

- Запуск локального веб-сервера на порту `8080`
- Додавання нового предмету через `POST` запит
- Перегляд списку всіх предметів (`GET /`)
- Перегляд конкретного предмету по ID (`GET /view?id=1`)
- Перегляд статистики по оцінках (`GET /stats`)

## Приклад використання
<img width="337" height="246" alt="image" src="https://github.com/user-attachments/assets/3bb4d148-d539-4f07-aaf7-0f7b6a5e7f82" />

## Clone the Repository
```git
git clone https://github.com/Xiancel/PersonalCard.git
cd PersonalCard
```

## ⚠️ Примітка

Для тестування веб-запитів **рекомендується використовувати Insomnia, Postman або аналогічні інструменти**, оскільки програма працює з HTTP-запитами (GET та POST) без графічного інтерфейсу.

Тестування напряму через браузер можливо лише для `GET` маршрутів (наприклад: `/`, `/view?id=1`, `/stats`).

Для `POST` запитів (наприклад: `/add`) потрібно використовувати Insomnia або Postman з відповідними параметрами.
