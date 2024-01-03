var stateMap = {
  "1": {"code": "FR-28", "name": "Eure-et-Loir"},
  "2": {"code": "FR-29", "name": "Finistère"},
  "3": {"code": "FR-22", "name": "Côtes-d'Armor"},
  "4": {"code": "FR-23", "name": "Creuse"},
  "5": {"code": "FR-21", "name": "Côte-d'Or"},
  "6": {"code": "FR-26", "name": "Drôme"},
  "7": {"code": "FR-27", "name": "Eure"},
  "8": {"code": "FR-24", "name": "Dordogne"},
  "9": {"code": "FR-25", "name": "Doubs"},
  "10": {"code": "FR-MQ", "name": "Martinique"},
  "11": {"code": "FR-94", "name": "Val-de-Marne"},
  "12": {"code": "FR-93", "name": "Seine-Saint-Denis"},
  "13": {"code": "FR-92", "name": "Hauts-de-Seine"},
  "14": {"code": "FR-91", "name": "Essonne"},
  "15": {"code": "FR-90", "name": "Territoire de Belfort"},
  "16": {"code": "FR-17", "name": "Charente-Maritime"},
  "17": {"code": "FR-16", "name": "Charente"},
  "18": {"code": "FR-15", "name": "Cantal"},
  "19": {"code": "FR-14", "name": "Calvados"},
  "20": {"code": "FR-13", "name": "Bouches-du-Rhône"},
  "21": {"code": "FR-12", "name": "Aveyron"},
  "22": {"code": "FR-11", "name": "Aude"},
  "23": {"code": "FR-10", "name": "Aube"},
  "24": {"code": "FR-2B", "name": "Haute-Corse"},
  "25": {"code": "FR-2A", "name": "Corse-du-Sud"},
  "26": {"code": "FR-19", "name": "Corrèze"},
  "27": {"code": "FR-18", "name": "Cher"},
  "28": {"code": "FR-88", "name": "Vosges"},
  "29": {"code": "FR-89", "name": "Yonne"},
  "30": {"code": "FR-80", "name": "Somme"},
  "31": {"code": "FR-81", "name": "Tarn"},
  "32": {"code": "FR-82", "name": "Tarn-et-Garonne"},
  "33": {"code": "FR-83", "name": "Var"},
  "34": {"code": "FR-84", "name": "Vaucluse"},
  "35": {"code": "FR-85", "name": "Vendée"},
  "36": {"code": "FR-86", "name": "Vienne"},
  "37": {"code": "FR-87", "name": "Haute-Vienne"},
  "38": {"code": "FR-01", "name": "Ain"},
  "39": {"code": "FR-02", "name": "Aisne"},
  "40": {"code": "FR-03", "name": "Allier"},
  "41": {"code": "FR-04", "name": "Alpes-de-Haute-Provence"},
  "42": {"code": "FR-05", "name": "Hautes-Alpes"},
  "43": {"code": "FR-06", "name": "Alpes-Maritimes"},
  "44": {"code": "FR-07", "name": "Ardèche"},
  "45": {"code": "FR-08", "name": "Ardennes"},
  "46": {"code": "FR-09", "name": "Ariège"},
  "47": {"code": "FR-RE", "name": "La Réunion"},
  "48": {"code": "FR-75", "name": "Paris"},
  "49": {"code": "FR-74", "name": "Haute-Savoie"},
  "50": {"code": "FR-77", "name": "Seien-et-Marne"},
  "51": {"code": "FR-76", "name": "Seine-Maritime"},
  "52": {"code": "FR-71", "name": "Saône-et-Loire"},
  "53": {"code": "FR-70", "name": "Haute-Saône"},
  "54": {"code": "FR-73", "name": "Savoie"},
  "55": {"code": "FR-72", "name": "Sarthe"},
  "56": {"code": "FR-79", "name": "Deux-Sèvres"},
  "57": {"code": "FR-78", "name": "Yvelines"},
  "58": {"code": "FR-YT", "name": "Mayotte"},
  "59": {"code": "FR-66", "name": "Pyrénées-Orientales"},
  "60": {"code": "FR-67", "name": "Bas-Rhin"},
  "61": {"code": "FR-64", "name": "Pyrénées-Atlantiques"},
  "62": {"code": "FR-65", "name": "Hautes-Pyrénées"},
  "63": {"code": "FR-62", "name": "Pas-de-Calais"},
  "64": {"code": "FR-63", "name": "Puy-de-Dôme"},
  "65": {"code": "FR-60", "name": "Oise"},
  "66": {"code": "FR-61", "name": "Orne"},
  "67": {"code": "FR-68", "name": "Haute-Rhin"},
  "68": {"code": "FR-69", "name": "Rhône"},
  "69": {"code": "FR-53", "name": "Mayenne"},
  "70": {"code": "FR-52", "name": "Haute-Marne"},
  "71": {"code": "FR-51", "name": "Marne"},
  "72": {"code": "FR-50", "name": "Manche"},
  "73": {"code": "FR-57", "name": "Moselle"},
  "74": {"code": "FR-56", "name": "Morbihan"},
  "75": {"code": "FR-55", "name": "Meuse"},
  "76": {"code": "FR-54", "name": "Meurhe-et-Moselle"},
  "77": {"code": "FR-59", "name": "Nord"},
  "78": {"code": "FR-58", "name": "Nièvre"},
  "79": {"code": "FR-48", "name": "Lozère"},
  "80": {"code": "FR-49", "name": "Maine-et-Loire"},
  "81": {"code": "FR-44", "name": "Loire-Atlantique"},
  "82": {"code": "FR-45", "name": "Loiret"},
  "83": {"code": "FR-46", "name": "Lot"},
  "84": {"code": "FR-47", "name": "Lot-et-Garonne"},
  "85": {"code": "FR-40", "name": "Landes"},
  "86": {"code": "FR-41", "name": "Loir-et-Cher"},
  "87": {"code": "FR-42", "name": "Loire"},
  "88": {"code": "FR-43", "name": "Haute-Loire"},
  "89": {"code": "FR-95", "name": "Val-d'Oise"},
  "90": {"code": "FR-GF", "name": "Guyane française"},
  "91": {"code": "FR-GP", "name": "Guadeloupe"},
  "92": {"code": "FR-39", "name": "Jura"},
  "93": {"code": "FR-38", "name": "Isère"},
  "94": {"code": "FR-31", "name": "Haute-Garonne"},
  "95": {"code": "FR-30", "name": "Gard"},
  "96": {"code": "FR-33", "name": "Gironde"},
  "97": {"code": "FR-32", "name": "Gers"},
  "98": {"code": "FR-35", "name": "Ille-et-Vilaine"},
  "99": {"code": "FR-34", "name": "Hérault"},
  "100": {"code": "FR-37", "name": "Indre-et-Loire"},
  "101": {"code": "FR-36", "name": "Indre"}
}
