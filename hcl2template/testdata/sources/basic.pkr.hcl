// a source represents a reusable setting for a system boot/start.
source "virtualbox-iso" "ubuntu-1204" {
    string   = "string"
    int      = 42
    int64    = 43
    bool     = true
    trilean  = true
    duration = "10s"
    map_string_string {
        a = "b"
        c = "d"
    }
    slice_string = [
        "a",
        "b",
        "c",
    ]

    // nested {
    //     string = "string"
    // }
}

// source "amazon-ebs" "ubuntu-1604" {
//     String =   "string"
//     Int =      42
//     Int64 =    43
//     Bool =     true
//     Trilean =  config.TriTrue
//     Duration = 10 * time.Second
//     MapStringString = map[string]string
//         "a" = "b"
//         "c" = "d"
//     },
//     SliceString = []string
//         "a",
//         "b",
//         "c",
//     },
// }

// source "amazon-ebs" "that-ubuntu-1.0" {
//     String =   "string"
//     Int =      42
//     Int64 =    43
//     Bool =     true
//     Trilean =  config.TriTrue
//     Duration = 10 * time.Second
//     MapStringString = map[string]string
//         "a" = "b"
//         "c" = "d"
//     },
//     SliceString = []string
//         "a",
//         "b",
//         "c",
//     },
// }
