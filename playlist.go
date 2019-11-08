package main

import "fmt"

// song in a playlist
type song struct {
	name       string
	artist     string
	songwriter string
	album      string
	genre      string
	next       *song
	previous   *song
}

// playlist containing songs
type playlist struct {
	name        string
	description string
	author      string
	head        *song
	tail        *song
	currentsong *song
}

// createPlaylist creates a playlist
func createPlaylist(name, description, author string) *playlist {
	return &playlist{
		name:        name,
		description: description,
		author:      author,
	}
}

//addSong adds song to playlist created
func (p *playlist) addSong(s song) *song {
	Song := &s
	if p.head == nil {
		p.head = Song
	} else {
		presentnode := p.tail
		presentnode.next = Song
		Song.previous = p.tail
	}
	p.tail = Song
	return nil
}

//playSong plays the current song in playlist
func (p *playlist) playSong() *song {
	p.currentsong = p.head
	return p.currentsong
}

//playPreviousSongs plays previous song in playlist
func (p *playlist) playPreviousSong() *song {
	p.currentsong = p.currentsong.previous
	return p.currentsong
}

//playNextSong plays next song in playlist
func (p *playlist) playNextSong() *song {
	p.currentsong = p.currentsong.next
	return p.currentsong
}

//showAllSongs shows all songs in playlist
func (p *playlist) showAllSongs() error {
	presentnode := p.head
	if presentnode == nil {
		fmt.Println("Empty playlist")
		return nil
	}
	fmt.Printf("%+v\n", *presentnode)
	for presentnode.next != nil {
		presentnode = presentnode.next
		fmt.Printf("%+v\n", *presentnode)
	}
	return nil
}

func main() {
	//creating playlist
	Playlist := createPlaylist("Billboard Hot 5", "best songs from BB5", "Billboard")
	fmt.Println("Created playlist " + Playlist.name)
	//songs to be added to playlist
	songs := [5]song{song{name: "Bohemian Rhapsody", artist: "Queen", songwriter: "Freddie Mercury", album: "A night at the Opera", genre: "Proressive Rock/Pop"},
		song{name: "Hey Jude", artist: "The Beatles", songwriter: "John Lennon/Paul McCartney", album: "Non album single", genre: "Rock"},
		song{name: "Sparks Fly", artist: "Taylor Swift", songwriter: "Taylor Swift", album: "Speak Now", genre: "Country Pop"},
		song{name: "Rolling in the Deep", artist: "Adele", songwriter: "Adele/Paul", album: "21", genre: "Pop"},
		song{name: "There's nothing holding me back", artist: "Shawn Mendes", songwriter: "Shawn Mendes", album: "Illminate", genre: "Pop Rock"}}

	//adding songs to playlist
	fmt.Println("Adding songs to playlist")
	for _, element := range songs {
		Playlist.addSong(element)
	}
	//showing all songs in playlist
	fmt.Println("Songs in " + Playlist.name)
	Playlist.showAllSongs()
	//Playing a song
	fmt.Println("play song")
	Playlist.playSong()
	fmt.Printf("Playing the song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
	//Playing next songs
	Playlist.playNextSong()
	fmt.Printf("Playing the next song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
	Playlist.playNextSong()
	fmt.Printf("Playing the next song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
	Playlist.playNextSong()
	fmt.Printf("Playing the next song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
	fmt.Println("play previous song")
	//Playing previoys songs
	Playlist.playPreviousSong()
	fmt.Printf("Playing the previous song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
	fmt.Println("play previous song")
	Playlist.playPreviousSong()
	fmt.Printf("Playing the previous song : %s ,artist %s\n", Playlist.currentsong.name, Playlist.currentsong.artist)
}
