# frozen_string_literal: true
# typed: true

# A marking class for when methods return void.
# Should never appear in types directly.
class T::Private::Types::Void < T::Types::Base
  ERROR_MESSAGE = "Validation is being done on an `Void`. Please report to #dev-productivity."

  # The actual return value of `.void` methods.
  #
  # Uses `module VOID` because this gives it a readable name when someone
  # examines it in Pry or with `#inspect` like:
  #
  #     T::Private::Types::Void::VOID
  #
  module VOID
    freeze
  end

  # @override Base
  def name
    "<VOID>"
  end

  # @override Base
  def valid?(obj)
    raise ERROR_MESSAGE
  end

  # @override Base
  private def subtype_of_single?(other)
    raise ERROR_MESSAGE
  end
end
