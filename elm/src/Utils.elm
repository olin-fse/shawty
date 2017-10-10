module Utils exposing (..)

import Http
import Json.Decode as Decode
import Json.Encode exposing (..)

import Msgs exposing (Msg(..))

shortenUrl : String -> Cmd Msg
shortenUrl url =
    let
        request =
            Http.post "/generate" (Http.jsonBody (encodeUserUrl url)) decodeGeneratedUrl
    in
        Http.send ShowUrl request

encodeUserUrl : String -> Value
encodeUserUrl url =
    object
        [ ("Url", string url)
        ]

decodeGeneratedUrl : Decode.Decoder String
decodeGeneratedUrl =
    Decode.at ["Url"] Decode.string
