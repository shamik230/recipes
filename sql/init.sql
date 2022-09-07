create table if not exists "recipies" (
    id bigserial primary key,
    name text not null,
    description text not null,
    dificulty text default 'пиздец сложно' not null
);

create table if not exists "recipies_tags" (
    id bigint not null,
    tag text
);


insert into recipies (name, description)
values
('Рамен', 'Здесь крутой рецепт рамена из подручных ингредиентой'),
('Бичпачмак', 'Крутейшее блюдо из татарстана');


insert into recipies_tags(id, tag)
values
( 1, 'яйца'),
( 1, 'лапша'),
( 2, 'мясо'),
( 2, 'картошка'),
( 2, 'тесто')

