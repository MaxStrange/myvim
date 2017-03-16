# This file will install all the repos and all the plugins necessary for
# my vimrc to work.
# It WILL NOT override your current .vimrc. To do that, run this, then simply
# cp my .vimrc over yours.

# Pathogen
echo "Installing Pathogen..."
mkdir -p ~/.vim/autoload ~/.vim/bundle
curl -LSso ~/.vim/autoload/pathogen.vim https://tpo.pe/pathogen.vim

# Rainbow Parens
echo "Installing rainbow parentheses..."
git clone https://github.com/kien/rainbow_parentheses.vim.git ~/.vim/bundle/rainbow_parentheses

# Elixir
echo "Installing elixir plugin..."
git clone https://github.com/elixir-lang/vim-elixir.git ~/.vim/bundle/vim-elixir

# CtrlP
echo "Installing ctrlP..."
git clone https://github.com/ctrlpvim/ctrlp.vim.git ~/.vim/bundle/ctrlp

# Buffergator
echo "Installing buffergator..."
git clone https://github.com/jeetsukumaran/vim-buffergator.git ~/.vim/bundle/buffergator

# Git fugitive
echo "Installing git fugitive..."
git clone git://github.com/tpope/vim-fugitive.git ~/.vim/bundle/vim-fugitive

# Vim airline
echo "Installing airline..."
git clone git://github.com/powerline/fonts ~/fonts
cd ~/fonts
./install.sh
cd ~/
git clone https://github.com/vim-airline/vim-airline ~/.vim/bundle/vim-airline

echo "Done. You now need to change your terminal font to something for powershell."
