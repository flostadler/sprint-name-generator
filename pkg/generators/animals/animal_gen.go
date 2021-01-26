package animals

import (
	"github.com/flostadler/name-generator/pkg/generators/preverb"
	"github.com/flostadler/name-generator/pkg/randomizer"
	"strings"
)

var(
	animals = [...] string {
		"Aardvark",
		"Alligator",
		"Crocodile",
		"Alpaca",
		"Ant",
		"Antelope",
		"Ape",
		"Armadillo",
		"Donkey",
		"Baboon",
		"Badger",
		"Bat",
		"Bear",
		"Beaver",
		"Bee",
		"Beetle",
		"Buffalo",
		"Butterfly",
		"Camel",
		"Carabao",
		"Water",
		"Caribou",
		"Cat",
		"Cattle",
		"Cheetah",
		"Chicken",
		"Rooster",
		"Chimpanzee",
		"Chinchilla",
		"Cicada",
		"Clam",
		"Cockroach",
		"Cod",
		"Coyote",
		"Crab",
		"Cricket",
		"Crow",
		"Raven",
		"Deer",
		"Dinosaur",
		"Dog",
		"Dolphin",
		"Duck",
		"Eagle",
		"Echidna",
		"Eel",
		"Elephant",
		"Elk",
		"Ferret",
		"Fish",
		"Fly",
		"Fox",
		"Frog",
		"Toad",
		"Gerbil",
		"Giraffe",
		"Gnat",
		"Gnu",
		"Goat",
		"Goldfish",
		"Goose",
		"Gorilla",
		"Grasshopper",
		"Guinea Pig",
		"Hamster",
		"Hare",
		"Hedgehog",
		"Herring",
		"Hippopotamus",
		"Hornet",
		"Horse",
		"Hound",
		"Hyena",
		"Impala",
		"Jackal",
		"Jellyfish",
		"Kangaroo",
		"Wallaby",
		"Koala",
		"Leopard",
		"Lion",
		"Lizard",
		"Llama",
		"Locust",
		"Louse",
		"Macaw",
		"Mallard",
		"Mammoth",
		"Manatee",
		"Marten",
		"Mink",
		"Minnow",
		"Mole",
		"Monkey",
		"Moose",
		"Mosquito",
		"Mouse",
		"Rat",
		"Mule",
		"Muskrat",
		"Otter",
		"Ox",
		"Oyster",
		"Panda",
		"Pig",
		"Hog",
		"Swine",
		"Platypus",
		"Porcupine",
		"Prairie Dog",
		"Pug",
		"Rabbit",
		"Raccoon",
		"Reindeer",
		"Rhinoceros",
		"Salmon",
		"Sardine",
		"Scorpion",
		"Seal",
		"Sea Lion",
		"Serval",
		"Shark",
		"Sheep",
		"Skunk",
		"Snail",
		"Snake",
		"Spider",
		"Squirrel",
		"Swan",
		"Termite",
		"Tiger",
		"Trout",
		"Turtle",
		"Walrus",
		"Wasp",
		"Weasel",
		"Whale",
		"Wolf",
		"Wombat",
		"Woodchuck",
		"Worm",
		"Yak",
		"Yellowjacket",
		"Zebra",
	}
)

func Generate() string {
	return strings.Title(preverb.GetRandomPreverb()) + " " + randomizer.GetRandom(animals[:])
}
