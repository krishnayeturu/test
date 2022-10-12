-- Create table AdmissionPolicy
create table if not exists `AdmissionPolicy`(
    `Id` int unsigned not null auto_increment,
    `UUID` varchar(200) not null unique,
    `Name` varchar(255) not null unique,
    `Type` varchar(255) not null default 'CREDENTIAL',
    primary key (Id),
    unique key `IX_AdmissionPolicy_Name` (`Name`),
    unique key `IX_AdmissionPolicy_UUID` (`UUID`)
) engine=InnoDB auto_increment=1 default charset =utf8mb4 collate =utf8mb4_unicode_ci;

-- Create relation table AdmissionPolicy
create table if not exists `AdmissionPolicyStatement`
(
    `Id`         int unsigned not null auto_increment,
    `PolicyUuid`   varchar(200) not null,
    `Effect`     varchar (5) not null default 'DENY',
    `Principal`   varchar(200) not null,
    `Action`     varchar(100) not null,
    `ResourceId` varchar(150) not null,
    primary key (Id),
constraint `FK_AdmissionPolicy_PolicyId`
    foreign key (`PolicyUuid`)
        references `AdmissionPolicy` (`UUID`)
        on delete cascade,
    unique key `UX_APR_PrincipalActionResourceId` (`Principal`, `Action`, `ResourceId`),
    index `IX_AdmissionPolicyStatement_PolicyId_Principal` (`PolicyUuid`, `Principal`)
) engine=InnoDB auto_increment=1 default charset =utf8mb4 collate =utf8mb4_unicode_ci;