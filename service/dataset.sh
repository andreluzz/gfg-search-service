#!/bin/sh
# Script to index new documents to elasticsearch. Execute passing the elasticsearch server host and port url. 
# Example: dataset.sh http://localhost:9200

endpoint="$1/products/_doc/_bulk"
curl -X POST $endpoint -H 'Content-Type: application/json' -d'
{ "index": { "_id": "1" } }
{"title": "Superstar Fashion Sneakers", "brand": "Adidas", "price": 50.97, "stock": 12}
{ "index": { "_id": "2" } }
{"title": "Grand Court Sneaker", "brand": "Adidas", "price": 43.79, "stock": 18}
{ "index": { "_id": "3" } }
{"title": "Continental 80 Sneaker", "brand": "Adidas", "price": 70.33, "stock": 2}
{ "index": { "_id": "4" } }
{"title": "Lite Racer Adapt Running Shoe", "brand": "Adidas", "price": 64.95, "stock": 6}
{ "index": { "_id": "5" } }
{"title": "Revolution 4 Running Shoe", "brand": "Nike", "price": 44.99, "stock": 20}
{ "index": { "_id": "6" } }
{"title": "Jordan Nike Kids Air 1 Retro High BG Basketball Shoe", "brand": "Nike", "price": 109.95, "stock": 4}
{ "index": { "_id": "7" } }
{"title": "Air Max Infuriate 2 Mid Basketball Shoe", "brand": "Nike", "price": 106, "stock": 8}
{ "index": { "_id": "8" } }
{"title": "Air Zoom Pegasus 35 Running Shoe", "brand": "Nike", "price": 83.81, "stock": 1}
{ "index": { "_id": "9" } }
{"title": "Air Monarch Iv Cross Trainer", "brand": "Nike", "price": 48.75, "stock": 13}
{ "index": { "_id": "10" } }
{"title": "Charged Assert 8 Running Shoe", "brand": "Under Armor", "price": 64.99, "stock": 0}
{ "index": { "_id": "11" } }
{"title": "Lockdown 4 Basketball Shoe", "brand": "Under Armor", "price": 64.95, "stock": 10}
{ "index": { "_id": "12" } }
{"title": "Micro G Assert 7 Sneaker", "brand": "Under Armor", "price": 57.99, "stock": 5}
{ "index": { "_id": "13" } }
{"title": "Ripple Elevated Sneaker", "brand": "Under Armor", "price": 60.83, "stock": 2}
{ "index": { "_id": "14" } }
{"title": "Suede Classic Plus Sneakers", "brand": "Puma", "price": 58.82, "stock": 22}
{ "index": { "_id": "15" } }
{"title": "Powertech Blaze Met Nm Running Shoe", "brand": "Puma", "price": 59.90, "stock": 3}
{ "index": { "_id": "16" } }
{"title": "Legend Short Sleeve Tee", "brand": "Nike", "price": 20.01, "stock": 13}
{ "index": { "_id": "17" } }
{"title": "Therma Swoosh Training Hoodie", "brand": "Nike", "price": 39.92, "stock": 22}
{ "index": { "_id": "18" } }
{"title": "Sportwear Club Shorts", "brand": "Nike", "price": 25.98, "stock": 8}
{ "index": { "_id": "19" } }
{"title": "3-Stripes Tee", "brand": "Adidas", "price": 20.73, "stock": 6}
{ "index": { "_id": "20" } }
{"title": "Superstar Trackpants", "brand": "Adidas", "price": 31.98, "stock": 15}
{ "index": { "_id": "21" } }
{"title": "Ultimate 365 Short (2019 Model)", "brand": "Adidas", "price": 51.97, "stock": 12}
{ "index": { "_id": "22" } }
{"title": "Archive Life T-Shirt", "brand": "Puma", "price": 13.95, "stock": 35}
{ "index": { "_id": "23" } }
{"title": "A.c.e. Sweat Jacket", "brand": "Puma", "price": 38.65, "stock": 16}
{ "index": { "_id": "24" } }
{"title": "Arsenal Fc Training Pants Pro with Zipper", "brand": "Puma", "price": 60.23, "stock": 2}
{ "index": { "_id": "25" } }
{"title": "Big Logo Tee Peacoat/Puma SM", "brand": "Puma", "price": 24.95, "stock": 9}
{ "index": { "_id": "26" } }
{"title": "Tech Short Sleeve T-Shirt", "brand": "Under Armor", "price": 14.50, "stock": 10}
{ "index": { "_id": "27" } }
{"title": "Fleece Pullover Hoodie", "brand": "Under Armor", "price": 31.96, "stock": 18}
{ "index": { "_id": "28" } }
{"title": "UA Raid Pocketed Short", "brand": "Under Armor", "price": 22.30, "stock": 11}
{ "index": { "_id": "29" } }
{"title": "Stripe Tech 1/4 Zip Pullover", "brand": "Under Armor", "price": 39.43, "stock": 4}
{ "index": { "_id": "30" } }
{"title": "Performance Polo", "brand": "Under Armor", "price": 39.99, "stock": 23}
'
