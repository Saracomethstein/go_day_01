CC = go
FLAGS = -o build
READ_DB_MAIN = cmd/go_day_01/DBReader/readDB.go
COMPARE_DB_MAIN = cmd/go_day_01/DBCompare/compareDB.go
COMPARE_FS_MAIN = cmd/go_day_01/FSCompare/compareFS.go
DATADIR = data/

all: build_dir  readDB compareDB compareFS

build_dir:
	@mkdir build

readDB:
	@$(CC) build $(FLAGS) $(READ_DB_MAIN)

compareDB:
	@$(CC) build $(FLAGS) $(COMPARE_DB_MAIN)

compareFS:
	@$(CC) build $(FLAGS) $(COMPARE_FS_MAIN)

clean:
	@rm -rf build