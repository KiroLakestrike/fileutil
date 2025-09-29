package fileutil


// Delete will be used to delete files in the future

// Check for folowing errors: 
// File does ot exitst, path is a directory, invalid file name, insufficient permissins, path is a symlink

// The main purpose of this file, is to delete a given file at the specified path. 
// Can be used to get rid of temporary files for examples, suring atomic copying for example.
// Will return true and nil if everything worked well, or false and err, if there was an error.
