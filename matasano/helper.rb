# typed: true

# Helper holds the methods to help solve some matasano crypto challenges
class Helper
  extend T::Sig

  sig { params(str: String).returns(String) }
  def hex_to_b64(str)
    # fill out this logic later... want to first try out sorbet type-checker
    str
  end
end

h = Helper.new.hex_to_b64(12)
