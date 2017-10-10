module Msgs exposing (..)

import Http

type Msg
    = Shorten
    | ShowUrl (Result Http.Error String)
    | Change String
