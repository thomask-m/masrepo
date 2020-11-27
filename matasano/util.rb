# typed: true
require 'sorbet-runtime'
require 'base64'

B64_TABLE = %w[
  A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
  a b c d e f g h i j k l m n o p q r s t u v w x y z
  0 1 2 3 4 5 6 7 8 9 + /
].freeze

# Ruby implementations of matasano crypto challenge
class Util
  extend T::Sig

  sig { params(num: String, ensure_size_multiple: Integer).returns(String) }
  def hex_to_b2(num, ensure_size_multiple)
    n = num.to_i(16)
    n_base_two = n.to_s(2)
    offset = ensure_size_multiple - (n_base_two.length % ensure_size_multiple)
    offset.times do
      n_base_two = "0#{n_base_two}"
    end
    n_base_two
  end

  sig { params(num: String).returns(String) }
  def hex_to_b64(num)
    num_in_base_two = hex_to_b2(num, 6)
    num_chunks = num_in_base_two.scan(/.{6}/)
    res = ''
    num_chunks.each do |chunk|
      unless chunk.is_a? Array
        i = chunk.to_i(2)
        res = "#{res}#{B64_TABLE[i]}"
      end
    end
    res
  end
end

hex = '49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'
p Util.new.hex_to_b64(hex)
