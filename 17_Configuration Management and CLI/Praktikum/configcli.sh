d=$(date)

#buat direktori nama dengan tanggal 
mkdir "$1 at $d"

#membuat folder about_me,my_friends,my_system_info
mkdir "$1 at $d"/about_me "$1 at $d"/my_friends "$1 at $d"/my_system_info

#membuat sub folder about_me
mkdir "$1 at $d"/about_me/personal "$1 at $d"/about_me/professional

#membuat file di folder
touch "$1 at $d"/about_me/personal/facebook.txt "$1 at $d"/about_me/professional/linkedin.txt "$1 at $d"/my_friends/list_of_my_friends.txt "$1 at $d"/my_system_info/about_this_laptop.txt "$1 at $d"/my_system_info/internet_connection.txt

#isi facebook argumen 2

echo "https://www.facebook.com/$2" > "$1 at $d"/about_me/personal/facebook.txt

#isi linkedin argumen 3
echo "https://www.linkedin.com/in/$3" > "$1 at $d"/about_me/professional/linkedin.txt

#memasukkan list of friends

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$1 at $d"/my_friends/list_of_my_friends.txt

#memasukkan nama user uname -a ke about_this_laptop

echo "My_username : $(whoami)" > "$1 at $d"/my_system_info/about_this_laptop.txt "Host : $(uname -a): $d" > "$1 at $d"/my_system_info/about_this_laptop.txt

#memasukkan ping google ke internet_connection.txt

ping google.com -c 3 > "$1 at $d"/my_system_info/internet_connection.txt
