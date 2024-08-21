CREATE TABLE user_roles(
    id SERIAL NOT NULL,
    user_id varchar(255) NOT NULL,
    role_id integer NOT NULL,
    role_names character varying[],
    PRIMARY KEY(id),
    CONSTRAINT user_roles_role_id_fkey FOREIGN key(role_id) REFERENCES roles(id)
);
CREATE UNIQUE INDEX user_roles_user_id_role_id_key ON user_roles USING btree ("user_id","role_id");