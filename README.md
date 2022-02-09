# Data Mask

##### Project that aims to mask sensitive data so that we can show the data structures in the logs if necessary without exposing the data of the user who is requesting the operation

* [Getting started](#Getting-started)
* [Limitations](#Limitations)

#### Getting started

###### We can create two types of masks for the data (using the show structure and the mask). We will address each of them and put examples of how to use them below.

##### 1. The mask tag has the following masking possibilities.

 |Tag        |Description                                                                                            |
|:---------:|:------------------------------------------------------------------------------------------------------|
 |initial  |hides the initial characters of a word, if it is a phrase mask the initial words|
|middle    |hides the characters in the middle of a word, if it is a sentence masks the words that are in the middle of the sentence  |
|last       |hide the final characters of a word, if it is a phrase mask the final words        |
|email       |hides email information in the format vi********@gmail.com |
|struct     |used to mask data that are in the internal structure, if not specified it will not check the fields belonging to the structure|
|all     |hide all data|
|firstLetter     |hides the initial characters of a word  |
|lastLetter     |hides the final characters of a word  |

###### 2. Getting the instance to mask the data and apply it.

``` golang
package main

import (
	masker "github.com/maskdata/mask"
	"fmt"
)

type User struct {
    Name string `json:"name" mask:"last"`
    Address `json:"address"`
}

type Address struct {
    Street string `json:"street" mask:"all"`
}

func main() {
	m := masker.NewMask()
    u := User {
        Name: "Peter Silva",
        Address: Address{
            Street: "Marginal Pinheiros",
        },
    }
	masked := m.MaskData(user)
    fmt.Println(masked)
}
```

###### 3. The return of this call is a string with the structure already masked based on the tags applied. As we can see, the Address structure did not get the mask because it does not have the struct tag in its declaration within the User structure

```
Result:
{"name": "Peter *****", "address": {"street": "Marginal Pinheiros"}}
```


##### 4. The tag show has a slightly different concept from the one mentioned above. It only shows the data that has been mapped with the tag, so we can say that it works the opposite way to the mask tag.


 |Tag        |Description                                                                                            |
|:---------:|:------------------------------------------------------------------------------------------------------|
 |initial  |expose the initial characters of a word, if it is a sentence expose the initial words|
|middle    |expose the characters in the middle of a word, if it is a sentence expose the words that are in the middle of the sentence  |
|last       |expose the final characters of a word, if it is a sentence expose the final words   |
|email      |expose the email in format vi********@gmail.com               |
|all        |expose all data|
|firstLetter |expose the initial characters of a word  |
|lastLetter     | expose the final characters of a word     |


###### 5. Getting the instance to display the data and apply it.

``` golang
package main

import (
	masker "github.com/maskdata/mask"
)

type User struct {
    Name string `json:"name" show:last`
    LastName string `json:"name"`
     Address `json:"address"`
}

type Address struct {
    Street string `json:"street" show:"all"`
}

func main() {
	m := masker.NewShow()
    u := User {
        Name: "Peter Jhonathan",
        LastName: "Silva",
         Address: Address {
            Street: "Marginal Pinheiros",
        },
    }
	masked := m.ShowData(user)
    fmt.Println(masked)
}
```

###### 6. The return of this call is a string of the structure with the tags applied.

```
Result:
{"name": "***** Jhonathan", "lastName": "*****", "address": {"street": "Marginal Pinheiros"}}
```

###### 7. If there is an error in the masking, it will return an empty string

# Limitations
###### Currently the lib only supports masking simple string data, not knowing how to deal with slices, feel free to evolve