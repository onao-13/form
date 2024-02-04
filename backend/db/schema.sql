CREATE TABLE people_employment(
    id BIGSERIAL PRIMARY KEY,
    organization_name VARCHAR(200) NOT NULL,
    short_organization_name VARCHAR(100) NOT NULL,
    type_participant VARCHAR(50) NOT NULL,
    position_leader_organization VARCHAR(100) NOT NULL,
    fio_leader_organization VARCHAR(150) NOT NULL,
    position_responsible_person VARCHAR(100) NOT NULL,
    fio_responsible_person VARCHAR(150) NOT NULL,
    telephone_number VARCHAR(11) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL
);