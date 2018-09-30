# Wikipedia Text Corpus Project

A (work in progress) set of tools for creating text corpus's from
wikipedia dumps. Wikipedia dumps come as a single, compressed, xml
file, filled with pages. The pages are unrendered templates that can
reference other pages. The document comes with an index that tells you
where to look in the file for a particular page, but no other indexes
are provided. This means finding all the pages or sentances with a
particular word would require a linear scan over the entire text
corpus. The tools provided here will process a wikipedia dump into an
indexed dataset. They will also render templates so you can process
the clean, plain english text that a user viewing a page on wikipedia
would see. Lastly they provide an api for constructing subsets of
articles based on their contents. e.g. all the articles with the word
swimming in the sports category.

## Loading the data

### Decompression

- Go's standard bzip2 decompression implementation is slow
- I'm playing with making it concurrent and faster
- Like to get thoughput to ~100 Mb/s, makes total processing time ~4 minutes
- Disk and a fast network should be able to support that, but maybe not.

### Parsing into pages

### Loading into a database

## Indexing and Querying

## Template resolution
