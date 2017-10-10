module Update exposing (..)

import Msgs exposing (Msg(..))
import Models exposing (Model)
import Utils exposing (..)

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        Shorten ->
            (model, shortenUrl model.userUrl)

        Change newContent ->
            ({ model | userUrl = newContent }, Cmd.none)

        ShowUrl (Ok newUrl) ->
            ({ model | generatedUrl = newUrl }, Cmd.none)

        ShowUrl (Err _) ->
            (model, Cmd.none)

