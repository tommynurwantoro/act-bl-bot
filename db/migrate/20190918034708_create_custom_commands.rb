class CreateCustomCommands < ActiveRecord::Migration[5.2]
  def change
    create_table :custom_commands, id: false, primary_key: :id do |t|
      t.primary_key :id, :unsigned_integer, auto_increment: true
      t.integer :group_id, null: false, limit: 8
      t.string :command, null: false
      t.string :message, null: false

      t.timestamps null: false
    end
  end
end
