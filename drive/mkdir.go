package drive

import (
    "google.golang.org/api/drive/v3"
    "io"
    "fmt"
)

const DirectoryMimeType = "application/vnd.google-apps.folder"

type MkdirArgs struct {
    Out io.Writer
    Name string
    Parents []string
    Share bool
}

func (self *Drive) Mkdir(args MkdirArgs) error {
    f, err := self.mkdir(args)
    if err != nil {
        return err
    }
    fmt.Printf("Directory '%s' created\n", f.Name)
    return nil
}

func (self *Drive) mkdir(args MkdirArgs) (*drive.File, error) {
    dstFile := &drive.File{Name: args.Name, MimeType: DirectoryMimeType}

    // Set parent folders
    dstFile.Parents = args.Parents

    // Create directory
    f, err := self.service.Files.Create(dstFile).Do()
    if err != nil {
        return nil, fmt.Errorf("Failed to create directory: %s", err)
    }

    fmt.Fprintf(args.Out, "\n[directory] id: %s, name: %s\n", f.Id, f.Name)

    //if args.Share {
    //    self.share(TODO)
    //}

    return f, nil
}
