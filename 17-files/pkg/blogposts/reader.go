package blogposts

import "io/fs"

func NewPostsFromFS(fileSystem fs.ReadDirFS) ([]Post, error) {
	posts := make([]Post, 0)

	dir, readDirErr := fileSystem.ReadDir(".")
	if readDirErr != nil {
		return posts, readDirErr
	}
	for _, entry := range dir {
		post, readPostErr := getPost(fileSystem, entry.Name())
		if readPostErr != nil {
			return posts, readPostErr
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, openErr := fileSystem.Open(fileName)
	if openErr != nil {
		return Post{}, openErr
	}
	defer postFile.Close()

	return newPost(postFile)
}
