class Tunacan < Formula
  desc ""
  homepage "https://github.com/hidexir/go-prompt-git"
  url "https://github.com/hidexir/go-prompt-git/releases/download/v0.1.5/go-prompt-git_0.1.5_Darwin_x86_64.tar.gz"
  version "0.1.5"
  sha256 "8811ccbf13381a07321d66b777c972e646a1f57eb8c053f573ddd786acfa3262"

  def install
    bin.install "gogit"
  end
end
