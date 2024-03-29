package hp

import (
	"github.com/flostadler/name-generator/pkg/randomizer"
)

var(
	characters = [...] string {
		"Aberforth Dumbledore",
		"Alastor (Mad-Eye) Moody",
		"Albert Runcorn",
		"Albus Dumbledore",
		"Albus Severus Potter",
		"Alecto Carrow",
		"Alice and Frank Longbottom",
		"Alicia Spinnet",
		"Amelia Bones",
		"Amos Diggory",
		"Amycus Carrow",
		"Andromeda Tonks",
		"Angelina Johnson",
		"Anthony Goldstein",
		"Antioch, Cadmus, and Ignotus Peverell",
		"Antonin Dolohov",
		"Arabella Figg",
		"Aragog",
		"Argus Filch",
		"Ariana Dumbledore",
		"Arthur Weasley",
		"Astoria Greengrass",
		"Augusta Longbottom",
		"Augustus Rookwood",
		"Aurora Sinistra",
		"Bane",
		"Barty Crouch Jr",
		"Barty Crouch Sr",
		"Bathilda Bagshot",
		"Beedle the Bard",
		"Bellatrix Lestrange",
		"Bertha Jorkins",
		"Bill Weasley",
		"Blaise Zabini",
		"Bob Ogden",
		"Bogrod",
		"Buckbeak",
		"Cedric Diggory",
		"Charity Burbage",
		"Charlie Weasley",
		"Cho Chang",
		"Colin Creevey",
		"Corban Yaxley",
		"Cormac McLaggen",
		"Cornelius Fudge",
		"Crabbe",
		"Crookshanks",
		"Cuthbert Binns",
		"Dean Thomas",
		"Dedalus Diggle",
		"Delphi Riddle",
		"Demelza Robins",
		"Dennis Creevey",
		"Dirk Cresswell",
		"Dobby",
		"Dolores Umbridge",
		"Draco Malfoy",
		"Dudley Dursley",
		"Eloise Midgen",
		"Elphias Doge",
		"Emmeline Vance",
		"Ernie Macmillan",
		"Errol",
		"Fang",
		"Fawkes",
		"Fenrir Greyback",
		"Filius Flitwick",
		"Firenze",
		"Fleur Delacour",
		"Fluffy",
		"Frank Bryce",
		"Fred Weasley",
		"Gabrielle Delacour",
		"Garrick Ollivander",
		"Gellert Grindelwald",
		"George Weasley",
		"Gilderoy Lockhart",
		"Ginny Weasley",
		"Godric Gryffindor",
		"Gornuk",
		"Goyle",
		"Graham Montague",
		"Grawp",
		"Great Aunt Muriel",
		"Gregorovitch",
		"Gregory Goyle",
		"Griphook",
		"Griselda Marchbanks",
		"Hannah Abbott",
		"Harry Potter",
		"Hedwig",
		"Helena Ravenclaw/The Grey Lady",
		"Helga Hufflepuff",
		"Hermione Granger",
		"Hokey",
		"Horace Slughorn",
		"Hugo Weasley",
		"Igor Karkaroff",
		"Irma Pince",
		"James Potter",
		"James Sirius Potter",
		"John Dawlish",
		"Justin Finch-Fletchley",
		"Katie Bell",
		"Kendra Dumbledore",
		"Kingsley Shacklebolt",
		"Kreacher",
		"Lavender Brown",
		"Lee Jordan",
		"Lily Luna Potter",
		"Lily Potter",
		"Lord Voldemort (Tom Marvolo Riddle)",
		"Lucius Malfoy",
		"Ludo Bagman",
		"Luna Lovegood",
		"Madam Malkin",
		"Madam Rosmerta",
		"Mafalda Hopkirk",
		"Magorian",
		"Marge Dursley",
		"Marietta Edgecombe",
		"Marvolo Gaunt",
		"Mary Riddle",
		"Merope Gaunt",
		"Millicent Bulstrode",
		"Minerva McGonagall",
		"Molly Weasley",
		"Morfin Gaunt",
		"Mundungus Fletcher",
		"Myrtle Warren/Moaning Myrtle",
		"Nagini",
		"Narcissa Malfoy",
		"Nearly Headless Nick",
		"Neville Longbottom",
		"Newt Scamander",
		"Nicolas Flamel",
		"Norbert",
		"Nott Sr.",
		"Nymphadora Tonks",
		"Oliver Wood",
		"Olympe Maxime",
		"Padma Patil",
		"Pansy Parkinson",
		"Parvati Patil",
		"Peeves",
		"Penelope Clearwater",
		"Percival Dumbledore",
		"Percy Weasley",
		"Peter Pettigrew",
		"Petunia Dursley",
		"Phineas Nigellus Black",
		"Pigwidgeon",
		"Pius Thicknesse",
		"Pomona Sprout",
		"Poppy Pomfrey",
		"Quirinus Quirrell",
		"Reginald Cattermole",
		"Remus Lupin",
		"Rita Skeeter",
		"Rolanda Hooch",
		"Romilda Vane",
		"Ron Weasley",
		"Ronan",
		"Rose Weasley",
		"Rowena Ravenclaw",
		"Rubeus Hagrid",
		"Rufus Scrimgeour",
		"Salazar Slytherin",
		"Scabbers",
		"Scabior",
		"Scorpius Malfoy",
		"Seamus Finnigan",
		"Septima Vector",
		"Severus Snape",
		"Silvanus Kettleburn",
		"Sir Cadogan",
		"Sirius Black",
		"Stan Shunpike",
		"Sturgis Podmore",
		"Susan Bones",
		"Sybill Trelawney",
		"Ted Tonks",
		"Terry Boot",
		"The Bloody Baron",
		"The Fat Friar",
		"The Fat Lady",
		"Theodore Nott",
		"Thomas Marvolo Riddle",
		"Thomas Riddle Jr.",
		"Thomas Riddle Sr.",
		"Thorfinn Rowle",
		"Travers",
		"Trevor",
		"Vernon Dursley",
		"Viktor Krum",
		"Vincent Crabbe",
		"Walden Macnair",
		"Wilhelmina Grubbly-Plank",
		"Wilkie Twycross",
		"Winky",
		"Xenophilius Lovegood",
		"Zacharias Smith",
	}
)

func Generate() string {
	return randomizer.GetRandom(characters[:])
}

func GenerateMultiple(count int) []string {
	s := make([]string, count)
	for i := 0;  i <= count; i++ {
		s[i] = Generate()
	}

	return s
}
