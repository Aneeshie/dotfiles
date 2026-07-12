{ config, pkgs, ... }:

let
  dotfiles = "${config.home.homeDirectory}/.dotfiles";
in
{
  home.username = "nara";
  home.homeDirectory = "/Users/nara";
  home.stateVersion = "26.05";

  programs.home-manager.enable = true;

  home.packages = with pkgs; [
    ripgrep
    fd
    fzf
    jq
    lazygit
    neovim
    tree-sitter
    nerd-fonts.iosevka-term
    tmux
  ];

  fonts.fontconfig.enable = true;

  home.sessionVariables = {
    EDITOR = "nvim";
  };

  programs.git = {
    enable = true;

    settings = {
      user = {
        name = "Aneeshie";
        email = "aneeshdas556@gmail.com";
      };
    };
  };

  programs.zsh = {
    enable = true;

    autosuggestion.enable = true;
    syntaxHighlighting.enable = true;

    initContent = ''
      bindkey '^f' autosuggest-accept
    '';

    shellAliases = {
      ".." = "cd ..";
      add = "git add .";
      push = "git push";
      pull = "git pull";
      n = "nvim";
      m = "git switch main";
    };
  };

  programs.starship = {
    enable = true;

    settings = {
      add_newline = false;
      format = "$directory$git_branch$git_status$cmd_duration$line_break$character";

      character = {
        success_symbol = "[>](purple)";
        error_symbol = "[>](red)";
      };

      cmd_duration = {
        format = "[$duration]($style) ";
      };
    };
  };


  # THIS IS FOR WEZTERM IF U WANT TO USE IT UNCOMMENT

  #home.file.".config/wezterm".source =
  #  config.lib.file.mkOutOfStoreSymlink
  #    "${dotfiles}/home/.config/wezterm";


   home.file.".config/herdr".source =
     config.lib.file.mkOutOfStoreSymlink
       "${dotfiles}/home/.config/herdr";


   home.file.".config/nvim".source =
     config.lib.file.mkOutOfStoreSymlink
       "${dotfiles}/home/.config/nvim";

  home.file.".config/aerospace".source =
    config.lib.file.mkOutOfStoreSymlink
      "${dotfiles}/home/.config/aerospace";

  home.file.".config/ghostty".source =
    config.lib.file.mkOutOfStoreSymlink
      "${dotfiles}/home/.config/ghostty";
}
