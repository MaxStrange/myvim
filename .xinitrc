[[ -f ~/.Xresources ]] && xrdb -merge ~/.Xresources
xset +fp /usr/share/fonts/local
xset fp rehash
xset -b
compton &
sh ~/.fehbg &

exec i3
