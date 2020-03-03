package repository

/**
start - Tentang bot ini
help - Nampilin semua perintah yang ada
halo - Cuma buat nyapa aja
titip_review - {title#url#telegram-users} Titip review PR
antrian_review - Nampilin semua antrian PR yang belum direview
sudah_direview - {urutan} Ngubah antrian review untuk yang sudah direview
sudah_direview_semua - {urutan} Ngubah antrian review untuk yang sudah direview untuk semua user
tambah_user_review - {urutan#telegram-users} Nambahin user ke antrian review
hapus_command - {urutan} Hapus antrian review
siap_qa - Mindahin antrian review ke antrian QA
antrian_qa - Nampilin semua antrian PR yang belum dites
sudah_dites - {urutan} Ngubah antrian QA untuk yang sudah dites
simpan_command - {kata#pesan} Kalau ada pengingat dengan perintah tertentu bisa pakai command ini loh
list_command - List custom command yang ada di group kamu
ubah_command - {urutan#pesan} Ubah isi pengingat yang ada di list command
hapus_command - {urutan} Hapus isi pengingat yang ada di list command
**/

// Command _
type Command struct {
	Name        string
	Description string
}

func GetCommand() *Command {
	return &Command{}
}

func (c *Command) All() []Command {
	// AllCommands List all commands
	return []Command{
		c.Start(),
		c.Help(),
		c.Halo(),
		c.TitipReview(),
		c.AntrianReview(),
		c.SudahDireview(),
		c.SudahDireviewSemua(),
		c.TambahUserReview(),
		c.HapusCommand(),
		c.SiapQA(),
		c.AntrianQA(),
		c.SudahDites(),
		c.SimpanCommand(),
		c.ListCommand(),
		c.UbahCommand(),
		c.HapusCommand(),
	}
}

func (c *Command) Start() Command {
	return Command{"/start", "Tentang bot ini"}
}

func (c *Command) Help() Command {
	return Command{"/help", "Nampilin semua perintah yang ada"}
}

func (c *Command) Halo() Command {
	return Command{"/halo", "Cuma buat nyapa aja"}
}

func (c *Command) TitipReview() Command {
	return Command{"/titip_review", "{title#url#telegram-users} Titip review PR"}
}

func (c *Command) AntrianReview() Command {
	return Command{"/antrian_review", "Nampilin semua antrian PR yang belum direview"}
}

func (c *Command) SudahDireview() Command {
	return Command{"/sudah_direview", "{urutan} Ngubah antrian review untuk yang sudah direview"}
}

func (c *Command) SudahDireviewSemua() Command {
	return Command{"/sudah_direview_semua", "{urutan} Ngubah antrian review untuk yang sudah direview untuk semua user"}
}

func (c *Command) TambahUserReview() Command {
	return Command{"/tambah_user_review", "{urutan#telegram-users} Nambahin user ke antrian review"}
}

func (c *Command) HapusReview() Command {
	return Command{"/hapus_review", "{urutan} Hapus antrian review"}
}

func (c *Command) SiapQA() Command {
	return Command{"/siap_qa", "{urutan} Mindahin antrian review ke antrian QA"}
}

func (c *Command) AntrianQA() Command {
	return Command{"/antrian_qa", "Nampilin semua antrian PR yang belum dites"}
}

func (c *Command) SudahDites() Command {
	return Command{"/sudah_dites", "{urutan} Ngubah antrian QA untuk yang sudah dites"}
}

func (c *Command) SendChat() Command {
	return Command{"/kirim_pesan", "{chatID#Messages} Kirim pesan ke chat yang dipilih"}
}

func (c *Command) SimpanCommand() Command {
	return Command{"/simpan_command", "{kata#pesan} Kalau ada pengingat dengan perintah tertentu bisa pakai command ini loh"}
}

func (c *Command) ListCommand() Command {
	return Command{"/list_command", "List custom command yang ada di group kamu"}
}

func (c *Command) UbahCommand() Command {
	return Command{"/ubah_command", "{urutan#pesan} Ubah isi pengingat yang ada di list command"}
}

func (c *Command) HapusCommand() Command {
	return Command{"/hapus_command", "{urutan} Hapus isi pengingat yang ada di list command"}
}
