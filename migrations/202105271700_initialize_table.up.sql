create table if not exists wagers
(
    id                    bigint unsigned auto_increment
        primary key,
    total_wager_value     int unsigned not null,
    odds                  int unsigned not null,
    selling_percentage    int not null,
    selling_price         decimal(10,2) not null,
    current_selling_price decimal(10,2) not null
        constraint wagers_current_selling_price_check
            check (current_selling_price >= 0),
    percentage_sold       decimal(10,2) null,
    amount_sold           decimal(10,2) null,
    placed_at             timestamp null
);


create table if not exists purchases
(
    id           int unsigned auto_increment
        primary key,
    wager_id     int unsigned not null,
    buying_price decimal(10,2) not null,
    bought_at    timestamp null
);



