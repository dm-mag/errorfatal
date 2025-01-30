# errorfatal
errorfatal package for golang

    This package gets rid of the annoying pattern:
    if e != nil {

    }

    and replaces it with:
    import e "github.com/dm-mag/errorfatal"
    ....
    e.F(os.Chdir("123"))
    ....
    e.F(os.WriteFile("123.txt", e.F2(os.ReadFile("456.txt")), 0644))
