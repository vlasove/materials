package mangrep

import "errors"

var (
	errNegativeAfterParam  = errors.New("mangrep: amount of AFTER can not be negative")
	errNegativeBeforeParam = errors.New("mangrep: amount of BEFORE can not be negative")
	errNegativeSidesParam  = errors.New("mangrep: amount of CONTEXT can not be negative")
)

type Options struct {
	amountAfter      int
	amountBefore     int
	amountSides      int
	isFixedTemplate  bool
	isNumLinePrinted bool
	isInversion      bool
	isIgnoreCase     bool
	isCounter        bool
}

func GetDefaultOptions() Options {
	return Options{
		isFixedTemplate:  false,
		isNumLinePrinted: false,
		isInversion:      false,
		isIgnoreCase:     false,
		isCounter:        false,
		amountAfter:      0,
		amountBefore:     0,
		amountSides:      0,
	}
}

var DefaultOptions = GetDefaultOptions()

type Option func(*Options) error

func SetIgnoreCaseOption(flag bool) Option {
	return func(o *Options) error {
		o.isIgnoreCase = flag
		return nil
	}
}

func SetCounterOption(flag bool) Option {
	return func(o *Options) error {
		o.isCounter = flag
		return nil
	}
}
func SetFixedTemplateOption(flag bool) Option {
	return func(o *Options) error {
		o.isFixedTemplate = flag
		return nil
	}
}

func SetNumLinePrintedOption(flag bool) Option {
	return func(o *Options) error {
		o.isNumLinePrinted = flag
		return nil
	}
}

func SetInversionOption(flag bool) Option {
	return func(o *Options) error {
		o.isInversion = flag
		return nil
	}
}

func SetAmountLinesPrintedAfterOption(amount int) Option {
	return func(o *Options) error {
		if amount < 0 {
			return errNegativeAfterParam
		}
		o.amountAfter = amount
		return nil
	}
}

func SetAmountLinesPrintedBeforeOption(amount int) Option {
	return func(o *Options) error {
		if amount < 0 {
			return errNegativeBeforeParam
		}
		o.amountBefore = amount
		return nil
	}
}

func SetAmountLinesPrintedSidesOption(amount int) Option {
	return func(o *Options) error {
		if amount < 0 {
			return errNegativeSidesParam
		}
		o.amountSides = amount
		return nil
	}
}
