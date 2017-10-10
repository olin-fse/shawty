import Html exposing (program)

import Models exposing (Model)
import Msgs exposing (Msg(..))
import Update exposing (update)
import View exposing (view)
import Subs exposing (subscriptions)

main : Program Never Model Msg
main =
    Html.program
        { init = init ""
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

init : String -> (Model, Cmd Msg)
init userUrl =
    ( Model userUrl ""
    , Cmd.none
    )
