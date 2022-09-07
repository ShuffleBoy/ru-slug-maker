package ruslug

import "testing"

func Test_toLower(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{
				source: "HeLlO WoRlD",
			},
			want: "hello world",
		},
		{
			name: "with russian chars",
			args: args{
				source: "ПРИВЕТ мир!",
			},
			want: "привет мир!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLower(tt.args.source); got != tt.want {
				t.Errorf("toLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "left",
			args: args{source: "  привет mir"},
			want: "привет mir",
		},
		{
			name: "right",
			args: args{source: "привет mir "},
			want: "привет mir",
		},
		{
			name: "both",
			args: args{source: " привет mir  "},
			want: "привет mir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trim(tt.args.source); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_escapeUnCompatibleChars(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "English text",
			args: args{
				source: "i'm thinking about her, and i will never forget her",
			},
			want: "im thinking about her and i will never forget her",
		},
		{
			name: "Russian text",
			args: args{
				source: "панацея — мифическое универсальное средство от всех болезней, способное также продлевать жизнь, вплоть до бесконечности",
			},
			want: "панацея  мифическое универсальное средство от всех болезней способное также продлевать жизнь вплоть до бесконечности",
		},
		{
			name: "Text with digits",
			args: args{
				source: "100 и 1 ночь",
			},
			want: "100 и 1 ночь",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeUnCompatibleChars(tt.args.source); got != tt.want {
				t.Errorf("escapeUnCompatibleChars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeWideSpaces(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single space",
			args: args{source: "privet mir"},
			want: "privet mir",
		},
		{
			name: "double space",
			args: args{source: "privet  mir"},
			want: "privet mir",
		},
		{
			name: "(I don't know how many)'le space",
			args: args{source: "privet              mir"},
			want: "privet mir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeWideSpaces(tt.args.source); got != tt.want {
				t.Errorf("removeWideSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transliterate(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple text",
			args: args{source: "амигос"},
			want: "amigos",
		},
		{
			name: "simple text",
			args: args{source: "во владивостоке столбы стоят не долго их сносят марки"},
			want: "vo vladivostoke stolby stoiat ne dolgo ikh snosiat marki",
		},
		{
			name: "combined text",
			args: args{source: "Mercedes-Benz торговая марка и одноимённая компания производитель легковых автомобилей премиального класса"},
			want: "Mercedes-Benz torgovaia marka i odnoimennaia kompaniia proizvoditel legkovykh avtomobilei premialnogo klassa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transliterate(tt.args.source); got != tt.want {
				t.Errorf("transliterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillWithDashes(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple text",
			args: args{source: "privet mir"},
			want: "privet-mir",
		},
		{
			name: "larger text",
			args: args{source: "сrazy fredrick bought many very exquisite opal jewels"},
			want: "сrazy-fredrick-bought-many-very-exquisite-opal-jewels",
		},
		{
			name: "text with digits",
			args: args{source: "vo vladivostoke 1 stolby stoiat ne dolgo ikh snosiat mark byvaet po 2 raza"},
			want: "vo-vladivostoke-1-stolby-stoiat-ne-dolgo-ikh-snosiat-mark-byvaet-po-2-raza",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fillWithDashes(tt.args.source); got != tt.want {
				t.Errorf("fillWithDashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{source: "Kafka как интеграционная платформа: от источников данных к потребителям и в хранилище"},
			want: "kafka-kak-integratsionnaia-platforma-ot-istochnikov-dannykh-k-potrebiteliam-i-v-khranilishche",
		},
		{
			name: "",
			args: args{source: "Секретики Unity3d. Зачем нужен флаг STARTER_ASSETS_PACKAGES_CHECKED в стартовых ассетахе"},
			want: "sekretiki-unity3d-zachem-nuzhen-flag-starterassetspackageschecked-v-startovykh-assetakhe",
		},
		{
			name: "",
			args: args{source: "PHP: атрибуты vs аннотации: оптимизируем метадату Doctrine"},
			want: "php-atributy-vs-annotatsii-optimiziruem-metadatu-doctrine",
		},
		{
			name: "",
			args: args{source: "Готовим Android к пентесту — WSA edition"},
			want: "gotovim-android-k-pentestu-wsa-edition",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.source); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
