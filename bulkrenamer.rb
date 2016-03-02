class Bulkrenamer < Formula
  desc "BulkRenamer"
  homepage "https://github.com/gosyujin/homebrew-bulkrenamer"
  head "https://github.com/gosyujin/homebrew-bulkrenamer.git"
  version "v0.1"
  url "#{homepage}/releases/download/#{version}/bulkrenamer"
  sha1 "5e3ac6469395aae60a00c61f45b5343207cca330"

  def install
    bin.install Dir['bulkrenamer']
  end
end
