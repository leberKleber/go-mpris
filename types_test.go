package mpris

import (
	"fmt"
	"testing"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
)

func TestMetadata_MPRISTrackID(t *testing.T) {
	var expectedTrackID dbus.ObjectPath
	var expectedErrorText string
	var trackID dbus.ObjectPath
	var err error

	//happycase
	expectedTrackID = "/my/path"
	expectedErrorText = "<nil>"
	trackID, err = Metadata{
		"mpris:trackid": dbus.MakeVariant(dbus.ObjectPath("/my/path")),
	}.MPRISTrackID()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackID, trackID, "unexpected trackID")

	//not present
	expectedTrackID = ""
	expectedErrorText = "<nil>"
	trackID, err = Metadata{}.MPRISTrackID()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackID, trackID, "unexpected trackID")

	//unexpected type
	expectedTrackID = ""
	expectedErrorText = "int could not be parsed to dbus.ObjectPath: the given type is not as expected"
	trackID, err = Metadata{
		"mpris:trackid": dbus.MakeVariant(123456789),
	}.MPRISTrackID()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackID, trackID, "unexpected trackID")
}

func TestMetadata_MPRISLength(t *testing.T) {
	var expectedLength int64
	var expectedErrorText string
	var length int64
	var err error

	//happycase
	expectedLength = 42
	expectedErrorText = "<nil>"
	length, err = Metadata{
		"mpris:length": dbus.MakeVariant(int64(42)),
	}.MPRISLength()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLength, length, "unexpected length")

	//not present
	expectedLength = 0
	expectedErrorText = "<nil>"
	length, err = Metadata{}.MPRISLength()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLength, length, "unexpected length")

	//unexpected type
	expectedLength = 0
	expectedErrorText = "string could not be parsed to int64: the given type is not as expected"
	length, err = Metadata{
		"mpris:length": dbus.MakeVariant("nope"),
	}.MPRISLength()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLength, length, "unexpected length")
}

func TestMetadata_MPRISArtURL(t *testing.T) {
	var expectedArtURL string
	var expectedErrorText string
	var artURL string
	var err error

	//happycase
	expectedArtURL = "/my/url"
	expectedErrorText = "<nil>"
	artURL, err = Metadata{
		"mpris:artUrl": dbus.MakeVariant("/my/url"),
	}.MPRISArtURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtURL, artURL, "unexpected artURL")

	//not present
	expectedArtURL = ""
	expectedErrorText = "<nil>"
	artURL, err = Metadata{}.MPRISArtURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtURL, artURL, "unexpected artURL")

	//unexpected type
	expectedArtURL = ""
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	artURL, err = Metadata{
		"mpris:artUrl": dbus.MakeVariant(42),
	}.MPRISArtURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtURL, artURL, "unexpected artURL")
}

func TestMetadata_XESAMAlbum(t *testing.T) {
	var expectedAlbum string
	var expectedErrorText string
	var album string
	var err error

	//happycase
	expectedAlbum = "/my/url"
	expectedErrorText = "<nil>"
	album, err = Metadata{
		"xesam:album": dbus.MakeVariant("/my/url"),
	}.XESAMAlbum()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbum, album, "unexpected album")

	//not present
	expectedAlbum = ""
	expectedErrorText = "<nil>"
	album, err = Metadata{}.XESAMAlbum()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbum, album, "unexpected album")

	//unexpected type
	expectedAlbum = ""
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	album, err = Metadata{
		"xesam:album": dbus.MakeVariant(42),
	}.XESAMAlbum()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbum, album, "unexpected album")
}

func TestMetadata_XESAMAlbumArtist(t *testing.T) {
	var expectedAlbumArtist []string
	var expectedErrorText string
	var albumArtist []string
	var err error

	//happycase
	expectedAlbumArtist = []string{"artist1", "artist2"}
	expectedErrorText = "<nil>"
	albumArtist, err = Metadata{
		"xesam:albumArtist": dbus.MakeVariant([]string{"artist1", "artist2"}),
	}.XESAMAlbumArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbumArtist, albumArtist, "unexpected albumArtist")

	//not present
	expectedAlbumArtist = nil
	expectedErrorText = "<nil>"
	albumArtist, err = Metadata{}.XESAMAlbumArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbumArtist, albumArtist, "unexpected albumArtist")

	//unexpected type
	expectedAlbumArtist = nil
	expectedErrorText = "int could not be parsed to []string: the given type is not as expected"
	albumArtist, err = Metadata{
		"xesam:albumArtist": dbus.MakeVariant(42),
	}.XESAMAlbumArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAlbumArtist, albumArtist, "unexpected albumArtist")
}

func TestMetadata_XESAMArtist(t *testing.T) {
	var expectedArtist []string
	var expectedErrorText string
	var artist []string
	var err error

	//happycase
	expectedArtist = []string{"artist1", "artist2"}
	expectedErrorText = "<nil>"
	artist, err = Metadata{
		"xesam:artist": dbus.MakeVariant([]string{"artist1", "artist2"}),
	}.XESAMArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtist, artist, "unexpected artist")

	//not present
	expectedArtist = nil
	expectedErrorText = "<nil>"
	artist, err = Metadata{}.XESAMArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtist, artist, "unexpected artist")

	//unexpected type
	expectedArtist = nil
	expectedErrorText = "int could not be parsed to []string: the given type is not as expected"
	artist, err = Metadata{
		"xesam:artist": dbus.MakeVariant(42),
	}.XESAMArtist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedArtist, artist, "unexpected artist")
}

func TestMetadata_XESAMAsText(t *testing.T) {
	var expectedAsText string
	var expectedErrorText string
	var asText string
	var err error

	//happycase
	expectedAsText = "asText"
	expectedErrorText = "<nil>"
	asText, err = Metadata{
		"xesam:asText": dbus.MakeVariant("asText"),
	}.XESAMAsText()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAsText, asText, "unexpected asText")

	//not present
	expectedAsText = ""
	expectedErrorText = "<nil>"
	asText, err = Metadata{}.XESAMAsText()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAsText, asText, "unexpected asText")

	//unexpected type
	expectedAsText = ""
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	asText, err = Metadata{
		"xesam:asText": dbus.MakeVariant(42),
	}.XESAMAsText()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAsText, asText, "unexpected asText")
}

func TestMetadata_XESAMAudioBPM(t *testing.T) {
	var expectedAudioBPM int
	var expectedErrorText string
	var audioBPM int
	var err error

	//happycase
	expectedAudioBPM = 4711
	expectedErrorText = "<nil>"
	audioBPM, err = Metadata{
		"xesam:audioBPM": dbus.MakeVariant(4711),
	}.XESAMAudioBPM()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAudioBPM, audioBPM, "unexpected audioBPM")

	//not present
	expectedAudioBPM = 0
	expectedErrorText = "<nil>"
	audioBPM, err = Metadata{}.XESAMAudioBPM()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAudioBPM, audioBPM, "unexpected audioBPM")

	//unexpected type
	expectedAudioBPM = 0
	expectedErrorText = "string could not be parsed to int: the given type is not as expected"
	audioBPM, err = Metadata{
		"xesam:audioBPM": dbus.MakeVariant("string"),
	}.XESAMAudioBPM()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAudioBPM, audioBPM, "unexpected audioBPM")
}

func TestMetadata_XESAMAutoRating(t *testing.T) {
	var expectedAutoRating float64
	var expectedErrorText string
	var autoRating float64
	var err error

	//happycase
	expectedAutoRating = 4711
	expectedErrorText = "<nil>"
	autoRating, err = Metadata{
		"xesam:autoRating": dbus.MakeVariant(float64(4711)),
	}.XESAMAutoRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAutoRating, autoRating, "unexpected autoRating")

	//not present
	expectedAutoRating = 0
	expectedErrorText = "<nil>"
	autoRating, err = Metadata{}.XESAMAutoRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAutoRating, autoRating, "unexpected autoRating")

	//unexpected type
	expectedAutoRating = 0
	expectedErrorText = "string could not be parsed to float64: the given type is not as expected"
	autoRating, err = Metadata{
		"xesam:autoRating": dbus.MakeVariant("string"),
	}.XESAMAutoRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedAutoRating, autoRating, "unexpected autoRating")
}

func TestMetadata_XESAMComment(t *testing.T) {
	var expectedComment []string
	var expectedErrorText string
	var comment []string
	var err error

	//happycase
	expectedComment = []string{"comment1", "comment2"}
	expectedErrorText = "<nil>"
	comment, err = Metadata{
		"xesam:comment": dbus.MakeVariant([]string{"comment1", "comment2"}),
	}.XESAMComment()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")

	//not present
	expectedComment = nil
	expectedErrorText = "<nil>"
	comment, err = Metadata{}.XESAMComment()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")

	//unexpected type
	expectedComment = nil
	expectedErrorText = "string could not be parsed to []string: the given type is not as expected"
	comment, err = Metadata{
		"xesam:comment": dbus.MakeVariant("string"),
	}.XESAMComment()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")
}

func TestMetadata_XESAMComposer(t *testing.T) {
	var expectedComment []string
	var expectedErrorText string
	var comment []string
	var err error

	//happycase
	expectedComment = []string{"comment1", "comment2"}
	expectedErrorText = "<nil>"
	comment, err = Metadata{
		"xesam:composer": dbus.MakeVariant([]string{"comment1", "comment2"}),
	}.XESAMComposer()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")

	//not present
	expectedComment = nil
	expectedErrorText = "<nil>"
	comment, err = Metadata{}.XESAMComposer()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")

	//unexpected type
	expectedComment = nil
	expectedErrorText = "string could not be parsed to []string: the given type is not as expected"
	comment, err = Metadata{
		"xesam:composer": dbus.MakeVariant("string"),
	}.XESAMComposer()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedComment, comment, "unexpected comment")
}

func TestMetadata_XESAMContentCreated(t *testing.T) {
	var expectedContentCreated time.Time
	var expectedErrorText string
	var contentCreated time.Time
	var err error

	//happycase
	expectedContentCreated = time.Date(2007, 4, 29, 13, 56, 0, 0, time.UTC)
	expectedErrorText = "<nil>"
	contentCreated, err = Metadata{
		"xesam:contentCreated": dbus.MakeVariant("2007-04-29T13:56+00:00"),
	}.XESAMContentCreated()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedContentCreated.Equal(contentCreated), "unexpected contentCreated", expectedContentCreated, contentCreated)

	//not present
	expectedContentCreated = time.Time{}
	expectedErrorText = "<nil>"
	contentCreated, err = Metadata{}.XESAMContentCreated()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedContentCreated.Equal(contentCreated), "unexpected contentCreated", expectedContentCreated, contentCreated)

	//unexpected type
	expectedContentCreated = time.Time{}
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	contentCreated, err = Metadata{
		"xesam:contentCreated": dbus.MakeVariant(42),
	}.XESAMContentCreated()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedContentCreated.Equal(contentCreated), "unexpected contentCreated", expectedContentCreated, contentCreated)

	//unexpected date format
	expectedContentCreated = time.Time{}
	expectedErrorText = `cound not parse time: parsing time "not a date-time" as "2006-01-02T15:04-07:00": cannot parse "not a date-time" as "2006": the given type is not as expected`
	contentCreated, err = Metadata{
		"xesam:contentCreated": dbus.MakeVariant("not a date-time"),
	}.XESAMContentCreated()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedContentCreated.Equal(contentCreated), "unexpected contentCreated", expectedContentCreated, contentCreated)
}

func TestMetadata_XESAMDiscNumber(t *testing.T) {
	var expectedDiscNumber int
	var expectedErrorText string
	var discNumber int
	var err error

	//happycase
	expectedDiscNumber = 42
	expectedErrorText = "<nil>"
	discNumber, err = Metadata{
		"xesam:discNumber": dbus.MakeVariant(42),
	}.XESAMDiscNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedDiscNumber, discNumber, "unexpected discNumber")

	//not present
	expectedDiscNumber = 0
	expectedErrorText = "<nil>"
	discNumber, err = Metadata{}.XESAMDiscNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedDiscNumber, discNumber, "unexpected discNumber")

	//unexpected type
	expectedDiscNumber = 0
	expectedErrorText = "string could not be parsed to int: the given type is not as expected"
	discNumber, err = Metadata{
		"xesam:discNumber": dbus.MakeVariant("string"),
	}.XESAMDiscNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedDiscNumber, discNumber, "unexpected discNumber")
}

func TestMetadata_XESAMFirstUsed(t *testing.T) {
	var expectedFirstUsed time.Time
	var expectedErrorText string
	var firstUsed time.Time
	var err error

	//happycase
	expectedFirstUsed = time.Date(2007, 4, 29, 13, 56, 0, 0, time.FixedZone("", 0))
	expectedErrorText = "<nil>"
	firstUsed, err = Metadata{
		"xesam:firstUsed": dbus.MakeVariant("2007-04-29T13:56+00:00"),
	}.XESAMFirstUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedFirstUsed.Equal(firstUsed), "unexpected firstUsed", expectedFirstUsed, firstUsed)

	//not present
	expectedFirstUsed = time.Time{}
	expectedErrorText = "<nil>"
	firstUsed, err = Metadata{}.XESAMFirstUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedFirstUsed.Equal(firstUsed), "unexpected firstUsed", expectedFirstUsed, firstUsed)

	//unexpected type
	expectedFirstUsed = time.Time{}
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	firstUsed, err = Metadata{
		"xesam:firstUsed": dbus.MakeVariant(42),
	}.XESAMFirstUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedFirstUsed.Equal(firstUsed), "unexpected firstUsed", expectedFirstUsed, firstUsed)

	//unexpected date format
	expectedFirstUsed = time.Time{}
	expectedErrorText = `cound not parse time: parsing time "not a date-time" as "2006-01-02T15:04-07:00": cannot parse "not a date-time" as "2006": the given type is not as expected`
	firstUsed, err = Metadata{
		"xesam:firstUsed": dbus.MakeVariant("not a date-time"),
	}.XESAMFirstUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedFirstUsed.Equal(firstUsed), "unexpected firstUsed", expectedFirstUsed, firstUsed)
}

func TestMetadata_XESAMGenre(t *testing.T) {
	var expectedGenre []string
	var expectedErrorText string
	var genre []string
	var err error

	//happycase
	expectedGenre = []string{"genre1", "genre2"}
	expectedErrorText = "<nil>"
	genre, err = Metadata{
		"xesam:genre": dbus.MakeVariant([]string{"genre1", "genre2"}),
	}.XESAMGenre()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedGenre, genre, "unexpected genre")

	//not present
	expectedGenre = nil
	expectedErrorText = "<nil>"
	genre, err = Metadata{}.XESAMGenre()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedGenre, genre, "unexpected genre")

	//unexpected type
	expectedGenre = nil
	expectedErrorText = "int could not be parsed to []string: the given type is not as expected"
	genre, err = Metadata{
		"xesam:genre": dbus.MakeVariant(42),
	}.XESAMGenre()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedGenre, genre, "unexpected genre")
}

func TestMetadata_XESAMLastUsed(t *testing.T) {
	var expectedLastUsed time.Time
	var expectedErrorText string
	var lastUsed time.Time
	var err error

	//happycase
	expectedLastUsed = time.Date(2007, 4, 29, 13, 56, 0, 0, time.FixedZone("", 0))
	expectedErrorText = "<nil>"
	lastUsed, err = Metadata{
		"xesam:lastUsed": dbus.MakeVariant("2007-04-29T13:56+00:00"),
	}.XESAMLastUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedLastUsed.Equal(lastUsed), "unexpected lastUsed", expectedLastUsed, lastUsed)

	//not present
	expectedLastUsed = time.Time{}
	expectedErrorText = "<nil>"
	lastUsed, err = Metadata{}.XESAMLastUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedLastUsed.Equal(lastUsed), "unexpected lastUsed", expectedLastUsed, lastUsed)

	//unexpected type
	expectedLastUsed = time.Time{}
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	lastUsed, err = Metadata{
		"xesam:lastUsed": dbus.MakeVariant(42),
	}.XESAMLastUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedLastUsed.Equal(lastUsed), "unexpected lastUsed", expectedLastUsed, lastUsed)

	//unexpected date format
	expectedLastUsed = time.Time{}
	expectedErrorText = `cound not parse time: parsing time "not a date-time" as "2006-01-02T15:04-07:00": cannot parse "not a date-time" as "2006": the given type is not as expected`
	lastUsed, err = Metadata{
		"xesam:lastUsed": dbus.MakeVariant("not a date-time"),
	}.XESAMLastUsed()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.True(t, expectedLastUsed.Equal(lastUsed), "unexpected lastUsed", expectedLastUsed, lastUsed)
}

func TestMetadata_XESAMLyricist(t *testing.T) {
	var expectedLyricist []string
	var expectedErrorText string
	var lyricist []string
	var err error

	//happycase
	expectedLyricist = []string{"genre1", "genre2"}
	expectedErrorText = "<nil>"
	lyricist, err = Metadata{
		"xesam:lyricist": dbus.MakeVariant([]string{"genre1", "genre2"}),
	}.XESAMLyricist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLyricist, lyricist, "unexpected lyricist")

	//not present
	expectedLyricist = nil
	expectedErrorText = "<nil>"
	lyricist, err = Metadata{}.XESAMLyricist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLyricist, lyricist, "unexpected lyricist")

	//unexpected type
	expectedLyricist = nil
	expectedErrorText = "int could not be parsed to []string: the given type is not as expected"
	lyricist, err = Metadata{
		"xesam:lyricist": dbus.MakeVariant(42),
	}.XESAMLyricist()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedLyricist, lyricist, "unexpected lyricist")
}

func TestMetadata_XESAMTitle(t *testing.T) {
	var expectedTitle string
	var expectedErrorText string
	var title string
	var err error

	//happycase
	expectedTitle = "title"
	expectedErrorText = "<nil>"
	title, err = Metadata{
		"xesam:title": dbus.MakeVariant("title"),
	}.XESAMTitle()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTitle, title, "unexpected title")

	//not present
	expectedTitle = ""
	expectedErrorText = "<nil>"
	title, err = Metadata{}.XESAMTitle()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTitle, title, "unexpected title")

	//unexpected type
	expectedTitle = ""
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	title, err = Metadata{
		"xesam:title": dbus.MakeVariant(42),
	}.XESAMTitle()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTitle, title, "unexpected title")
}

func TestMetadata_XESAMTrackNumber(t *testing.T) {
	var expectedTrackNumber int
	var expectedErrorText string
	var trackNumber int
	var err error

	//happycase
	expectedTrackNumber = 42
	expectedErrorText = "<nil>"
	trackNumber, err = Metadata{
		"xesam:trackNumber": dbus.MakeVariant(42),
	}.XESAMTrackNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackNumber, trackNumber, "unexpected trackNumber")

	//not present
	expectedTrackNumber = 0
	expectedErrorText = "<nil>"
	trackNumber, err = Metadata{}.XESAMTrackNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackNumber, trackNumber, "unexpected trackNumber")

	//unexpected type
	expectedTrackNumber = 0
	expectedErrorText = "string could not be parsed to int: the given type is not as expected"
	trackNumber, err = Metadata{
		"xesam:trackNumber": dbus.MakeVariant("string"),
	}.XESAMTrackNumber()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedTrackNumber, trackNumber, "unexpected trackNumber")
}

func TestMetadata_XESAMURL(t *testing.T) {
	var expectedURL string
	var expectedErrorText string
	var url string
	var err error

	//happycase
	expectedURL = "url"
	expectedErrorText = "<nil>"
	url, err = Metadata{
		"xesam:url": dbus.MakeVariant("url"),
	}.XESAMURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedURL, url, "unexpected url")

	//not present
	expectedURL = ""
	expectedErrorText = "<nil>"
	url, err = Metadata{}.XESAMURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedURL, url, "unexpected url")

	//unexpected type
	expectedURL = ""
	expectedErrorText = "int could not be parsed to string: the given type is not as expected"
	url, err = Metadata{
		"xesam:url": dbus.MakeVariant(42),
	}.XESAMURL()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedURL, url, "unexpected url")
}

func TestMetadata_XESAMUseCount(t *testing.T) {
	var expectedUseCount int
	var expectedErrorText string
	var useCount int
	var err error

	//happycase
	expectedUseCount = 42
	expectedErrorText = "<nil>"
	useCount, err = Metadata{
		"xesam:useCount": dbus.MakeVariant(42),
	}.XESAMUseCount()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUseCount, useCount, "unexpected useCount")

	//not present
	expectedUseCount = 0
	expectedErrorText = "<nil>"
	useCount, err = Metadata{}.XESAMUseCount()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUseCount, useCount, "unexpected useCount")

	//unexpected type
	expectedUseCount = 0
	expectedErrorText = "string could not be parsed to int: the given type is not as expected"
	useCount, err = Metadata{
		"xesam:useCount": dbus.MakeVariant("string"),
	}.XESAMUseCount()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUseCount, useCount, "unexpected useCount")
}

func TestMetadata_XESAMUserRating(t *testing.T) {
	var expectedUserRating float64
	var expectedErrorText string
	var userRating float64
	var err error

	//happycase
	expectedUserRating = 4711
	expectedErrorText = "<nil>"
	userRating, err = Metadata{
		"xesam:userRating": dbus.MakeVariant(float64(4711)),
	}.XESAMUserRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUserRating, userRating, "unexpected userRating")

	//not present
	expectedUserRating = 0
	expectedErrorText = "<nil>"
	userRating, err = Metadata{}.XESAMUserRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUserRating, userRating, "unexpected userRating")

	//unexpected type
	expectedUserRating = 0
	expectedErrorText = "string could not be parsed to float64: the given type is not as expected"
	userRating, err = Metadata{
		"xesam:userRating": dbus.MakeVariant("string"),
	}.XESAMUserRating()
	assert.Equal(t, expectedErrorText, fmt.Sprint(err), "unexpected error text")
	assert.Equal(t, expectedUserRating, userRating, "unexpected userRating")
}

func TestMetadata_Find(t *testing.T) {
	expectedValue := dbus.MakeVariant("123456")

	value, found := Metadata{
		"myKey": dbus.MakeVariant("123456"),
	}.Find("myKey")
	assert.True(t, found)
	assert.Equal(t, expectedValue, value)

}
