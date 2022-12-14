package repository

import (
	model "github.com/AlexanderTurok/beat-store-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	Create(account model.Account) (int, error)
	GetId(email, password string) (int, error)
}

type Account interface {
	Confirm(username string) error
	Get(accountId int) (model.Account, error)
	Update(accountId int, input model.AccountUpdateInput) error
	GetPasswordHash(accountId int) (string, error)
	Delete(accountId int) error
}

type Artist interface {
	Create(accountId int) error
	Get(accountId int) (model.Artist, error)
	GetAll() ([]model.Artist, error)
	GetPasswordHash(accountId int) (string, error)
	Delete(accountId int) error
}

type Payment interface {
	CreatePaymentAccount(accountId int, stripeId string) error
}

type Product interface {
	Create(artistId int, stripeId string) (int64, error)
}

type Beat interface {
	Create(productId int64, input model.Beat) (int, error)
	Get(beatId int) (model.Beat, error)
	GetAll() ([]model.Beat, error)
	GetArtistsBeat(artistId, beatId int) (model.Beat, error)
	GetAllArtistsBeats(artistId int) ([]model.Beat, error)
	Update(beatId int, input model.BeatUpdateInput) error
	Delete(beatId int) error
}

type Playlist interface {
	Create(accountId int, input model.Playlist) (int, error)
	Get(playlistId int) (model.Playlist, error)
	GetAll() ([]model.Playlist, error)
	GetAccountsPlaylist(accountId, playlistId int) (model.Playlist, error)
	GetAllAccountsPlaylists(accountId int) ([]model.Playlist, error)
	Update(playlistId int, input model.PlaylistUpdateInput) error
	Delete(playlistId int) error
	AddBeat(playlistId, beatId int) error
	GetBeat(playlistId, beatId int) (model.Beat, error)
	GetAllBeats(playlistId int) ([]model.Beat, error)
	GetBeatFromAccountsPlaylists(accountId, playlistId, beatId int) (model.Beat, error)
	GetAllBeatsFromAccountsPlaylists(accountId, playlistId int) ([]model.Beat, error)
	DeleteBeat(playlistId, beatId int) error
}

type Repositories struct {
	Auth     Auth
	Account  Account
	Artist   Artist
	Payment  Payment
	Product  Product
	Beat     Beat
	Playlist Playlist
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Auth:     NewAuthRepository(db),
		Account:  NewAccountRepository(db),
		Artist:   NewArtistRepository(db),
		Payment:  NewPaymentRepository(db),
		Product:  NewProductRepository(db),
		Beat:     NewBeatRepository(db),
		Playlist: NewPlaylistRepository(db),
	}
}
