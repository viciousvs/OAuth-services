CREATE TABLE IF NOT EXISTS stuff_base(
    id serial PRIMARY KEY NOT NULL,
    name varchar(50) NOT NULL,
    price money not null,
    description text not null
);
/* power integer not null CHECK power > 0,*/
create table IF NOT EXISTS fridge(
    id serial PRIMARY KEY not null,
    base_id int not null REFERENCES stuff_base(id),
    color varchar(50) not null,
    weight decimal not null CHECK (weight > 0),
    width decimal not null CHECK (width > 0),
    height decimal not null CHECK (height > 0),
    depth decimal not null CHECK (depth > 0),
    number_of_doors int not null CHECK (number_of_doors > 0),
    volume decimal not null CHECK (volume > 0),
    freezer_volume int not null CHECK (freezer_volume> 0),
    number_of_compressors int not null CHECK (number_of_compressors > 0),
    coolant varchar(50) not null, /*Хладагент*//**/
    power int not null check (power > 0),
    material varchar(30) not null
);

create table IF NOT EXISTS washer(
    id serial PRIMARY KEY not null,
    base_id int not null REFERENCES stuff_base(id),
    color varchar(50) not null,
    wash_type varchar(30) not null,
    weight decimal not null CHECK (weight > 0),
    width decimal not null CHECK (width > 0),
    height decimal not null CHECK (height > 0),
    depth decimal not null CHECK (depth > 0),
    max_load_weight int not null check (max_load_weight > 0),
    rotational_speed int not null check ( rotational_speed > 0 ),
    material varchar(30) not null
);

create table if not exists dishwasher(
    id serial primary key not null,
    base_id int not null REFERENCES stuff_base(id),
    color varchar(50) not null,
    weight decimal not null CHECK (weight > 0),
    width decimal not null CHECK (width > 0),
    height decimal not null CHECK (height > 0),
    depth decimal not null CHECK (depth > 0),
    water_consumption int not null check ( water_consumption > 0 ),
    material varchar(30) not null
);

create table if not exists hoover(
     id serial primary key not null,
     base_id int not null REFERENCES stuff_base(id),
     color varchar(50) not null,
     dust_collector decimal not null,
     weight decimal not null CHECK (weight > 0),
     power int not null check (power > 0),
     material varchar(30) not null,
     cord_length int not null CHECK (cord_length>0)
);

create table if not exists kettle(
     id serial primary key not null,
     base_id int not null REFERENCES stuff_base(id),
     power int not null check (power > 0),
     material varchar(30) not null,
     descaling_filter bool not null,
     weight decimal not null CHECK (weight > 0),
     volume decimal not null CHECK (volume > 0),
     color varchar(50) not null
);
create table if not exists toaster(
      id serial primary key not null,
      base_id int not null REFERENCES stuff_base(id),
      power int not null check (power > 0),
      material varchar(30) not null,
      branches int not null check (branches>0),
      toast_count int not null check (toaster.toast_count>0),
      color varchar(50) not null
);

insert into stuff_base (name, price, description)
values ('RB33A32N0SA', 262500, 'Cool fridge');
insert into fridge (base_id, color, weight, width, height, depth, number_of_doors, volume, freezer_volume,
                    number_of_compressors, coolant, power, material)
values (1, 'blue', 64.5, 59.5, 185, 67.5, 2, 328, 98, 1, 'R600a (изобутан)', 280, 'plastic/metal');


insert into stuff_base (name, price, description)
values ('MFN60-S1003', 136175, 'cool washer');
insert into washer (base_id, color, wash_type, weight, width, height, depth, max_load_weight, rotational_speed,
                    material)
values (2, 'blue', 'auto', 50, 59, 58, 40, 6, 1200, 'plastic/metal');


insert into stuff_base (name, price, description)
values ('DWB12-7718', 210450, 'cool dishwasher');
insert into dishwasher (base_id, color, weight, width, height, depth, water_consumption, material)
values (3, 'blue', 46, 59.8, 81.5, 55, 10, 'metal/glass/plastic');


insert into stuff_base (name, price, description)
values ('Vacuum T30', 191990, 'cool hoover for home');
insert into hoover (base_id, color, dust_collector, weight, power, material, cord_length)
values (4, 'blue', 0.6, 1.75, 2900, 'metal/plastic', 15);


insert into stuff_base (name, price, description)
values ('RK-G178', 13999, 'cool electronic kettle');
insert into kettle (base_id, power, material, descaling_filter, weight, volume, color)
values (5, 2000, 'plastic/glass', true, 1.0, 1.7, 'blue');


insert into stuff_base (name, price, description)
values ('DSL-601', 21990, 'cool toaster');
insert into toaster (base_id, power, material, branches, toast_count, color)
values (6, 650, 'plastic', 2, 2, 'blue');