module SignIn exposing (..)
import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL

type LoginStatus = Initial | Failure | Loading | Success
type alias Model
  = { status: LoginStatus, email:String, password: String 
      
  }


init : () -> (Model, Cmd Msg)
init _ = ({status = Initial, email = "", password = ""}, Cmd.none)



-- UPDATE


type Msg
  = UpdateEmail String| UpdatePassword String | Login
  | GotResponse (Result Http.Error String)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    UpdateEmail str-> ({model | email = str }, Cmd.none)
    UpdatePassword str -> ({model | password = str}, Cmd.none)
    Login ->
      ({model | status = Loading}, login)

    GotResponse result ->
      case result of
        Ok _ ->
          ({model | status = Success}, Cmd.none)

        Err _ ->
          ({model | status = Failure, password = ""}, Cmd.none)



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions _ =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  Html.form [ onSubmit Login ]
    [ label [] [ text "email" ]
    , input [value model.email, onInput UpdateEmail] []
    , label [] [ text "password" ]
    , input [ type_ "password", value model.password, onInput UpdatePassword] []
    , button [] [ text "login"]
    , viewResponse model
    ]


viewResponse : Model -> Html Msg
viewResponse model =
  case model.status of
    Failure ->
      span []
        [ text "Failed to sign in :(" ]

    Loading ->
      text "Loading..."

    Success -> span [] [text "Success!"]
    _ -> text ""

viewSignUp : Model -> Html Msg
viewSignUp model =
    Html.form [] [ 
        label [] [ text "Email" ]
        , input [] []
        , label [] [ text "Password"]
        , input [] []
        , button [] [ text "Sign Up" ]
    ]



-- HTTP


login : Cmd Msg
login =
  Http.post
    { url = "http://localhost:8080/api/v1/login"
    , expect = Http.expectString GotResponse 
    , body = Http.stringBody "text/plain" ""
    }
