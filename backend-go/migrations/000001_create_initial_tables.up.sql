CREATE TABLE nationality (
    nationality_id INT GENERATED ALWAYS AS identity primary key,
    nationality_name varchar(50) not null,
    nationality_code char(2) not null
);

CREATE TABLE customer (
    cst_id INT GENERATED ALWAYS AS identity primary key,
    nationality_id INT NOT NULL,
    cst_name CHAR(50) NOT NULL,
    cst_phoneNum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL,
    CONSTRAINT fk_nationality FOREIGN KEY(nationality_id) REFERENCES nationality(nationality_id)
);

CREATE TABLE family_list (
    fl_id INT GENERATED ALWAYS AS identity primary key,
    cst_id INT NOT NULL,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL,
    CONSTRAINT fk_customer FOREIGN KEY(cst_id) REFERENCES customer(cst_id)
);
