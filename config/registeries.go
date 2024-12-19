package config

type ArmorTrimMaterail struct {
	AssetName      string
	Ingredient     string
	ItemModelIndex float32
	Description    string
}

type ArmorTrimPattern struct {
	AssetId      string
	TemplateItem string
	Description  string
	Decal        bool
}

type BannerPattern struct {
	AssetId        string
	TranslationKey string
}

type Biome struct {
	HasPrecipitation   bool
	Tempreture         float32
	TempretureModifier string
	Downfall           float32
	Effect             Effects
}

type Effects struct {
	FogColor           int32
	WaterColor         int32
	WaterFogColor      int32
	SkyColor           int32
	FoliageColor       int32
	GrassColor         int32
	GrassColorModifier string
	Particle           Particle
	AmbientSound       string
	MoodSound          MoodSound
	AdditionsSound     AdditionsSound
	Music              Music
}

type Particle struct {
	Options     ParticleOptions
	Probability float32
}

type ParticleOptions struct {
	Type  string
	Value any
}

type MoodSound struct {
	Sound             string
	TickDelay         int32
	BlockSearchExtent int32
	Offset            float64
}

type AdditionsSound struct {
	Sound      string
	TickChance float64
}

type Music struct {
	Sound               string
	MinDelay            int32
	MaxDelay            int32
	ReplaceCurrentMusic bool
}

type ChatType struct {
	Chat      Decoration
	Narration Decoration
}

type Decoration struct {
	TranslationKey string
	Style          any
	Parameters     []string
}

type DamageType struct {
	MessageId        string
	Scaling          string
	Exhaustion       float32
	Effects          string
	DeathMessageType string
}

type DimensionType struct {
	FixedTime                   int64
	HasSkylight                 bool
	HasCeiling                  bool
	Ultrawarm                   bool
	Natural                     bool
	CordinateScale              float64
	BedWorks                    bool
	RespawnAnchorWorks          bool
	MinY                        int32
	Height                      int32
	LogicalHight                int32
	Infiniburn                  string
	Effects                     string
	AmbientLight                float32
	PiglinSafe                  bool
	HasRaids                    bool
	MonsterSpawnLightLevel      int32
	MonsterSpawnBlockLightLimit int32
}

type WolfVariant struct {
	WildTexture  string
	TameTexture  string
	AngryTexture string
	Biomes       []string
}

type PaintingVarint struct {
	AssetId string
	Height  int32
	Width   int32
}
