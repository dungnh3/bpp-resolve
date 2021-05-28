create table wagers
(
    id                    bigint unsigned auto_increment
        primary key,
    total_wager_value     int unsigned not null,
    odds                  int unsigned not null,
    selling_percentage    bigint not null,
    selling_price         double not null,
    current_selling_price double not null
        constraint wagers_current_selling_price_check
            check (current_selling_price >= 0),
    percentage_sold       bigint null,
    amount_sold           bigint null,
    placed_at             datetime(3) null
);


create table purchases
(
    id           int unsigned auto_increment
        primary key,
    wager_id     int unsigned not null,
    buying_price double not null,
    bought_at    datetime(3) null
);



