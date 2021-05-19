
exports.up = function(knex) {
    return knex.schema.createTable('pages', (table) => {
        table.increments();
        table.string('name')
        table.string('url');
        table.boolean('status')
    })
};

exports.down = function(knex) {

    return knex.schema.dropTable('pages');

};
