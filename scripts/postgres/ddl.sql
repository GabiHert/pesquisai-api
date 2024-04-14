create schema perguntai;



create table perguntai.request (

                                   id UUID primary key,

                                   context VARCHAR(1000),

                                   research VARCHAR(1000),

                                   total_researches INT default 0,

                                   total_finished_researched INT default 0,

                                   status VARCHAR(10),

                                   overall TEXT,

                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                                   updated_at TIMESTAMP

);



create table perguntai.location (

                                    id UUID primary key,

                                    name VARCHAR(10)

);



create table perguntai.request_location (

                                            request_id UUID,

                                            location_id UUID,

                                            primary key (request_id,location_id),

                                            foreign key (request_id) references request(id),

                                            foreign key (location_id) references location(id)

);



create table perguntai.language (

                                    id UUID primary key,

                                    name VARCHAR(10)

);



create table perguntai.request_language (

                                            request_id UUID,

                                            language_id UUID,

                                            primary key (request_id,language_id),

                                            foreign key (request_id) references request(id),

                                            foreign key (language_id) references language(id)

);



create table perguntai.research (

                                    id UUID primary key,

                                    request_id UUID ,

                                    title VARCHAR(50),

                                    link VARCHAR(100),

                                    status VARCHAR(10),

                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                                    updated_at TIMESTAMP,

                                    summary TEXT,

                                    foreign key (request_id) references request(id)

)