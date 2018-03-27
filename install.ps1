function install
{
    Param ([uri] $gitlink, [String] $dirname)
    $path = Join-Path C:\Users\maxst\vimfiles\bundle $dirname
    If (-Not (Test-Path $path))
    {
        Write-Host "Installing $dirname..."
        git clone $gitlink $path
    }
    Set-Location $path
    vim -c q "helptags doc/"
}

# Exit the script on cmdlet error
$ErrorActionPreference = "Stop"

If (-Not (Test-Path C:\Users\maxst\vimfiles\bundle))
{
    Write-Host "Installing Pathogen..."
    New-Item -Path C:\Users\maxst\vimfiles\autoload -Type Directory -Force
    New-Item -Path C:\Users\maxst\vimfiles\bundle -Type Directory
    Invoke-WebRequest -Uri https://tpo.pe/pathogen.vim -OutFile C:\Users\maxst\vimfiles\autoload\pathogen.vim
}

install https://github.com/kien/rainbow_parentheses.vim.git rainbow_parentheses
install https://github.com/elixir-lang/vim-elixir.git vim-elixir
install https://github.com/ctrlpvim/ctrlp.vim.git ctrlp
install https://github.com/jeetsukumaran/vim-buffergator.git buffergator
install https://github.com/tpope/vim-fugitive.git vim-fugitive
install https://github.com/tmhedberg/SimpylFold.git SimplyFold
install https://github.com/PProvost/vim-ps1.git vim-ps1
install https://github.com/rust-lang/rust.vim.git rust.vim

If (-Not (Test-Path C:\Users\maxst\vimfiles\ftdetect))
{
    Write-Host "Installing vim-scala..."
    New-Item -Type Directory -Path C:\Users\maxst\vimfiles\ftdetect
    New-Item -Type Directory -Path C:\Users\maxst\vimfiles\indent
    New-Item -Type Directory -Path C:\Users\maxst\vimfiles\syntax
    Invoke-WebRequest -Uri "https://raw.githubusercontent.com/derekwyatt/vim-scala/master/ftdetect/scala.vim" -OutFile "C:\Users\maxst\vimfiles\ftdetect\scala.vim"
    Invoke-WebRequest -Uri "https://raw.githubusercontent.com/derekwyatt/vim-scala/master/indent/scala.vim" -OutFile  "C:\Users\maxst\vimfiles\indent\scala.vim"
    Invoke-WebRequest -Uri "https://raw.githubusercontent.com/derekwyatt/vim-scala/master/syntax/scala.vim" -OutFile "C:\Users\maxst\vimfiles\syntax\scala.vim"
}

If (-Not (Test-Path C:\Users\maxst\vimfiles\bundle\vim-airline))
{
    If (-Not (Test-Path C:\Users\maxst\fonts))
    {
        Write-Host "Installing airline..."
        git clone git://github.com/powerline/fonts C:\Users\maxst\fonts
        Set-Location C:\Users\maxst\fonts
        .\install.ps1
    }
    Set-Location C:\Users\maxst
    git clone https://github.com/vim-airline/vim-airline C:\Users\maxst\vimfiles\bundle\vim-airline
}
Set-Location C:\Users\maxst\vimfiles\bundle\vim-airline
vim -c q "helptags doc/"
