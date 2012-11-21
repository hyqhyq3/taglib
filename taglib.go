package taglib

// #cgo LDFLAGS: -ltag_c
// #include <taglib/tag_c.h>
import "C"
import "errors"

type TagLib_File struct{
	*_Ctype_TagLib_File
}

type TagLib_Tag struct {
	*_Ctype_TagLib_Tag
}

var OpenFileError = errors.New("cannot open file")

var GetTagError = errors.New("cannot get tag")

func NewFile(filename string) (f *TagLib_File, err error) {
	cfilename := C.CString(filename)
	cfile := C.taglib_file_new(cfilename)
	if cfile == nil {
		err = OpenFileError
	}
	f = &TagLib_File{cfile}
	return
}

func (f *TagLib_File) Tag() (t *TagLib_Tag, err error) {
	ctag := C.taglib_file_tag(f._Ctype_TagLib_File)
	if ctag == nil {
		err = GetTagError
	}
	t = &TagLib_Tag{ctag}
	return
}

func (t *TagLib_Tag) Title() (title string) {
	ctitle := C.taglib_tag_title(t._Ctype_TagLib_Tag)
	title = C.GoString(ctitle)
	return
}

func(t *TagLib_Tag) Artist() (artist string) {
	cartist:= C.taglib_tag_artist(t._Ctype_TagLib_Tag)
	artist = C.GoString(cartist)
	return
}

func(t *TagLib_Tag) Album() (album string) {
	calbum:= C.taglib_tag_album(t._Ctype_TagLib_Tag)
	album = C.GoString(calbum)
	return
}


func(t *TagLib_Tag) Comment() (comment string) {
	ccomment:= C.taglib_tag_comment(t._Ctype_TagLib_Tag)
	comment = C.GoString(ccomment)
	return
}

func(t *TagLib_Tag) Genre() (genre string) {
	cgenre:= C.taglib_tag_genre(t._Ctype_TagLib_Tag)
	genre = C.GoString(cgenre)
	return
}

func(t *TagLib_Tag) Track() (track uint) {
	ctrack:= C.taglib_tag_track(t._Ctype_TagLib_Tag)
	track = uint(ctrack)
	return
}
